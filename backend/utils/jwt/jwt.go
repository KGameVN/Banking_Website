package jwt

import (
	"time"
	"fmt"
	"comb.com/banking/internal/errors"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("jwtglobal123kcikre-43569c939dgdfj3") // bạn nên lưu secret này qua biến môi trường

func GenerateJWT(password, username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"password": password,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // token sống 1 ngày
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseJWT(tokenStr string) (jwt.MapClaims, error) {
	if tokenStr == "" {
		return nil, &errors.AppError{Code: errors.ErrEmptyToken.Code,
			Message: errors.ErrEmptyToken.Message, Err: nil}
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &errors.AppError{Code: errors.ErrParseJWTToken.Code,
				Message: errors.ErrParseJWTToken.Message, Err: nil}
		}
		return jwtSecret, nil
	})

	if err != nil || token == nil {
		fmt.Println("JWT Parse Error:", err)
		return nil, &errors.AppError{Code: errors.ErrInvalidJWTToken.Code,
			Message: errors.ErrInvalidJWTToken.Message, Err: nil}
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return claims, &errors.AppError{Code: errors.ErrInvalidJWTToken.Code,
			// Message: errors.ErrEmptyToken.Message, Err: nil}
			Message: "asdf", Err: nil}
	}

	return nil, err
}

