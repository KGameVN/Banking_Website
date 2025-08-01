package services

import (
	"log"
	"net/http"

	"comb.com/banking/ent/user"
	"comb.com/banking/ent/useraccount"
	"comb.com/banking/errors"
	"github.com/labstack/echo/v4"
)

func (s *Service) GetAccountInfo(c echo.Context) error {
	userID := c.Get("userID").(int)

	ctx := c.Request().Context()

	account, err := s.Repository.DbClient.UserAccount.
		Query().
		Where(useraccount.HasUserWith(user.IDEQ(userID))).
		Only(ctx)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Không tìm thấy tài khoản",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"account_number": account.AccountNumber,
		"balance":        account.Balance,
	})
}

func (s Service) Transaction(c echo.Context) error {
	// userID := c.Param("id")

	type Request struct {
		Id     int64  `json:"id"`
		Amount int64  `json:"amount"`
		Type   string `json:"type"` // "deposit" hoặc "withdraw"
		Time   string `json:"time"` // time on transaction
	}
	var req Request
	if err := c.Bind(&req); err != nil {
		return &errors.AppError{Code: errors.ErrAccountNotFound.Code, Message: "ádfsadj", Err: err}
		// return c.JSON(http.StatusBadRequest, echo.Map{"error": "Dữ liệu không hợp lệ"})
	}

	ctx := c.Request().Context()

	account, err := s.Repository.DbClient.UserAccount.
		Query().
		Where(useraccount.HasUserWith(user.AccountNumber(1))).
		Only(ctx)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Không tìm thấy tài khoản"})
	}

	if req.Type != "dep" {
		req.Amount = -req.Amount
	}

	newBalance := account.Balance + req.Amount
	if newBalance < 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Số dư không đủ"})
	}

	_, err = s.Repository.DbClient.UserAccount.
		UpdateOneID(account.ID).
		SetBalance(newBalance).
		Save(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Không thể cập nhật số dư"})
	}

	return c.JSON(http.StatusOK, echo.Map{"balance": newBalance})
}

func (s Service) Transfer(c echo.Context) error {
	// userID := c.Get("userID").(int)

	type Request struct {
		FromAccountNumber int64 `json:"from_account_number"`
		ToAccountNumber   int   `json:"to_account_number"`
		Amount            int64 `json:"amount"`
	}

	var req Request
	if err := c.Bind(&req); err != nil || req.Amount <= 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Dữ liệu không hợp lệ"})
	}

	ctx := c.Request().Context()

	// ✅ Bắt đầu transaction
	tx, err := s.Repository.DbClient.Tx(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Không thể bắt đầu giao dịch"})
	}
	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
			panic(r)
		}
	}()

	// Tài khoản gửi
	fromAcc, err := tx.UserAccount.
		Query().
		Where(useraccount.HasUserWith(user.AccountNumber(200))).
		Only(ctx)
	if err != nil {
		_ = tx.Rollback()
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Không tìm thấy tài khoản gửi"})
	}

	// Tài khoản nhận
	toAcc, err := tx.UserAccount.
		Query().
		Where(useraccount.AccountNumberEQ(int64(req.ToAccountNumber))).
		Only(ctx)
	if err != nil {
		_ = tx.Rollback()
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Không tìm thấy tài khoản nhận"})
	}

	// Không thể chuyển cho chính mình
	if fromAcc.ID == toAcc.ID {
		_ = tx.Rollback()
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Không thể chuyển cho chính mình"})
	}

	// Kiểm tra số dư
	if fromAcc.Balance < req.Amount {
		_ = tx.Rollback()
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Số dư không đủ"})
	}

	// Trừ tiền người gửi
	_, err = tx.UserAccount.
		UpdateOneID(fromAcc.ID).
		SetBalance(fromAcc.Balance - req.Amount).
		Save(ctx)
	if err != nil {
		_ = tx.Rollback()
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Không thể trừ tiền"})
	}

	// Cộng tiền người nhận
	_, err = tx.UserAccount.
		UpdateOneID(toAcc.ID).
		SetBalance(toAcc.Balance + req.Amount).
		Save(ctx)
	if err != nil {
		_ = tx.Rollback()
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Không thể cộng tiền cho người nhận"})
	}

	// ✅ Commit transaction
	if err := tx.Commit(); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Không thể hoàn tất giao dịch"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Chuyển khoản thành công",
	})
}
