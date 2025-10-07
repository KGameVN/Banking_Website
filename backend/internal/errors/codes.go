// dinh nghia code (domaincode) + mapping http code

package errors

import "net/http"

type Code int

const (
	CodeUnknown Code = iota
	CodeInvalid = 100
	CodeUnauthorized = 101
	CodeForbidden = 102
	CodeNotFound = 103
	CodeConflict = 104
	CodeInsufficientFunds = 201
	CodeInternal = 500
)

// HTTPStatus maps domain Code -> HTTP status
func (c Code) HTTPStatus() int {
	switch c {
	case CodeInvalid:
		return http.StatusBadRequest
	case CodeUnauthorized:
		return http.StatusUnauthorized
	case CodeForbidden:
		return http.StatusForbidden
	case CodeNotFound:
		return http.StatusNotFound
	case CodeConflict:
		return http.StatusConflict
	case CodeInsufficientFunds:
		return http.StatusBadRequest
	case CodeInternal:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}


