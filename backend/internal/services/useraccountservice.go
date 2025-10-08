package services

import (
	"net/http"

	"comb.com/banking/api"
	"comb.com/banking/ent/transfer"
	"comb.com/banking/ent/user"
	"comb.com/banking/ent/useraccount"
	"comb.com/banking/internal/errors"
	"comb.com/banking/utils/converter"
	"github.com/labstack/echo/v4"
	// "github.com/segmentio/kafka-go"
)

func (s *Service) GetAccountInfo(c echo.Context) error {
	userAccount, err := converter.StringToInt64(c.Param("account"))
	if err != nil {
		return &errors.AppError{Code: errors.ErrIDIsNotValid.Code,
			Message: errors.ErrAccountNotFound.Message, Err: err}
	}

	ctx := c.Request().Context()

	account, err := s.Repository.DbClient.UserAccount.
		Query().
		Where(useraccount.HasUserWith(user.ID(int(userAccount)))).
		Only(ctx)

	if err != nil {
		return &errors.AppError{Code: errors.ErrAccountNotFound.Code,
			Message: errors.ErrAccountNotFound.Message, Err: err}
	}

	return api.Success(c, echo.Map{
		"user_account": account.Edges.User,
		"balance":      account.Balance,
	})

}

func (s Service) Transaction(c echo.Context) error {
	userAccount, err := converter.StringToInt64(c.Param("account"))
	if err != nil {
		return &errors.AppError{Code: errors.ErrIDIsNotValid.Code,
			Message: errors.ErrAccountNotFound.Message, Err: err}
	}

	type Request struct {
		Amount int64  `json:"amount"`
		Type   string `json:"type"` // "deposit" hoặc "withdraw"
		Time   string `json:"time"` // time on transaction
	}
	var req Request
	if err := c.Bind(&req); err != nil {
		return &errors.AppError{Code: errors.ErrInvalidJsonFormat.Code,
			Message: errors.ErrInvalidJsonFormat.Message, Err: err}
	}

	ctx := c.Request().Context()

	account, err := s.Repository.DbClient.UserAccount.
		Query().
		Where(useraccount.AccountNumber(userAccount)).
		Only(ctx)
	if err != nil {
		return &errors.AppError{Code: errors.ErrAccountNotFound.Code,
			Message: errors.ErrAccountNotFound.Message, Err: err}
	}

	if req.Type != "dep" {
		req.Amount = -req.Amount
	}

	newBalance := account.Balance + req.Amount
	if newBalance < 0 {
		return &errors.AppError{Code: errors.ErrNotEnoughBalance.Code,
			Message: errors.ErrNotEnoughBalance.Message, Err: err}
	}

	_, err = s.Repository.DbClient.UserAccount.
		UpdateOneID(account.ID).
		SetBalance(newBalance).
		Save(ctx)
	if err != nil {
		return &errors.AppError{Code: errors.ErrCanUpdateDB.Code,
			Message: errors.ErrCanUpdateDB.Message, Err: err}
	}

	return c.JSON(http.StatusOK, echo.Map{"balance": newBalance})
}

func (s Service) Transfer(c echo.Context) error {

	type Request struct {
		FromAccount int   `json:"from_account_number"`
		ToAccount   int   `json:"to_account_number"`
		Amount      int64 `json:"amount"`
	}

	var req Request
	if err := c.Bind(&req); err != nil || req.Amount <= 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Dữ liệu không hợp lệ"})
	}

	ctx := c.Request().Context()
	// 1. Tạo bản ghi transfer
	// transferRecord, err := s.Repository.DbClient.Transfer.Create().
	// 	SetFromAccountNumber(req.FromAccount).
	// 	SetToAccount(req.ToAccount).
	// 	SetBalance(req.Amount).
	// 	Save(ctx)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to create transfer: %w", err)
	// }

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
		Where(useraccount.AccountNumberEQ(int64(req.FromAccount))).
		Only(ctx)
	if err != nil {
		_ = tx.Rollback()
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Không tìm thấy tài khoản gửi"})
	}

	// Tài khoản nhận
	toAcc, err := tx.UserAccount.
		Query().
		Where(useraccount.AccountNumberEQ(int64(req.ToAccount))).
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

	return api.Success(c, echo.Map{
		"Source Account":      fromAcc,
		"Destination Account": toAcc,
		"Transferred Amount":  req.Amount,
		"Status":              "Success",
		"Transfer ID":         fromAcc.ID,
	})
}

func (s *Service) GetTransHistory(c echo.Context) error {
	userAccount, err := converter.StringToInt64(c.Param("account"))
	if err != nil {
		return &errors.AppError{Code: errors.ErrIDIsNotValid.Code,
			Message: errors.ErrAccountNotFound.Message, Err: err}
	}

	ctx := c.Request().Context()

	account, err := s.Repository.DbClient.UserAccount.
		Query().
		Where(useraccount.AccountNumberEQ(userAccount)).
		WithAccountNumberID(). // eager load transactions
		Only(ctx)
	if err != nil {
		return &errors.AppError{Code: errors.ErrAccountNotFound.Code,
			Message: errors.ErrAccountNotFound.Message, Err: err}
	}

	return api.Success(c, echo.Map{
		"transactions": account.Edges.AccountNumberID,
	})
}
