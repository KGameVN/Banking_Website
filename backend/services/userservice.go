package services

import (
	"net/http"

	"comb.com/banking/ent"
	"comb.com/banking/ent/user"
	"comb.com/banking/utils"
	"github.com/labstack/echo"
)

func (s Service) Login(c echo.Context) error {
	// Lấy dữ liệu JSON từ body
	body, err := utils.JsonToMap(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	username, ok1 := body["username"].(string)
	password, ok2 := body["password"].(string)
	if !ok1 || !ok2 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Missing email or password"})
	}

	// Truy vấn user trong DB
	foundUser, err := s.Repository.DbClient.User.
		Query().
		Where(
			user.UsernameEQ(username),
			user.PasswordEQ(password), // ❗ Nếu dùng hash: chỉ kiểm tra email, sau đó dùng bcrypt.CompareHashAndPassword
		).
		Only(c.Request().Context())

	if err != nil {
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid email or password"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Database error"})
	}

	// ✅ Nếu thành công
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Login successful",
		"user": echo.Map{
			"id":       foundUser.ID,
			"username": foundUser.Username,
			"email":    foundUser.Email,
			"fullname": foundUser.Fullname,
		},
	})
}
