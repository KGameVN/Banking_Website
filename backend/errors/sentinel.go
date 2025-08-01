// sentinel errors (common domain errors)

package errors

// Sentinel errors (reusable, testable)
var (
	ErrUserNotFound    = New(CodeNotFound, "user not found")
	ErrAccountNotFound = New(CodeNotFound, "account not found")
	ErrInsufficient    = New(CodeInsufficientFunds, "insufficient funds")
	ErrInvalidRequest  = New(CodeInvalid, "invalid request")
	
	// login errors
	ErrUnauthorized   = New(CodeUnauthorized, "user unauthorized")
	// login //JWT errors
	ErrEmptyToken     = New(CodeUnauthorized, "token is empty")
	ErrInvalidJWTToken = New(CodeUnauthorized, "invalid JWT token")
	ErrParseJWTToken   = New(CodeUnauthorized, "can't parse JWT token")
	ErrExpiredJWTToken = New(CodeUnauthorized, "expired JWT token")
)
