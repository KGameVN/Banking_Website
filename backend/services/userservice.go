package services

import (
	"log"
	"net/http"
	"time"

	"comb.com/banking/ent"
	"comb.com/banking/ent/token"
	"comb.com/banking/ent/user"
	"comb.com/banking/utils/converter"
	"comb.com/banking/utils/jwt"
	"github.com/labstack/echo/v4"
)

func (s Service) Login(c echo.Context) error {
	// Lấy dữ liệu JSON từ body
	body, err := converter.JsonToMap(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	username, ok1 := body["username"].(string)
	password, ok2 := body["password"].(string)
	rememberMe := body["rememberMe"].(bool)
	if !ok1 || !ok2 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Missing email or password"})
	}

	// Truy vấn user trong DB
	foundUser, err := s.Repository.DbClient.User.
		Query().
		Where(
			user.UsernameEQ(username),
			user.PasswordEQ(password), // ❗ Nếu dùng hash: chỉ kiểm tra email, sau đó dùng bcrypt.CompareHashAndPassword
		).WithTokens(
		func(tq *ent.TokenQuery) {
			tq.Where(
				token.ExpiredtimeGT(time.Now()), // expiredtime > now
				token.IsUsingEQ(true),           // is_using = true
			)
		}).
		Only(c.Request().Context())

	if err != nil {
		log.Println(err)
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid email or password"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Database error"})
	}

	if len(foundUser.Edges.Tokens) == 0 {
		// jwt
		token, err := jwt.GenerateJWT(foundUser.Username, password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Could not generate token"})
		}

		expiryDate := time.Now().Add(1 * time.Hour)
		if rememberMe {
			expiryDate = time.Now().Add(120 * time.Hour)
		}
		_, err = s.Repository.DbClient.Token.Create().
			SetUserID(foundUser.ID).
			SetToken(token).
			SetExpiredtime(expiryDate).
			SetType("login").
			Save(c.Request().Context())

		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
		}
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Login successful",
			"token":   token,
			"user": echo.Map{
				"id":            foundUser.ID,
				"accountnumber": foundUser.AccountNumber,
				"username":      foundUser.Username,
				"email":         foundUser.Email,
			},
		})
	}
	// response
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Login successful",
		"token":   foundUser.Edges.Tokens[0].Token,
		"user": echo.Map{
			"id":            foundUser.ID,
			"accountnumber": foundUser.AccountNumber,
			"username":      foundUser.Username,
			"email":         foundUser.Email,
		},
	})
}

func (s *Service) ProfileService(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "hello",
	})
}

func (s *Service) GetTransHistory(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetTransHistory",
	})
}

func (s *Service) Register(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "register",
	})
}
