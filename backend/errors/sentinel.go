// sentinel errors (common domain errors)

package errors

// Sentinel errors (reusable, testable)
var (
	ErrUserNotFound    = New(CodeNotFound, "User not found")
	ErrAccountNotFound = New(CodeNotFound, "Account not found")
	ErrInsufficient    = New(CodeInsufficientFunds, "Insufficient funds")
	ErrInvalidRequest  = New(CodeInvalid, "Invalid request")
	
	// login errors
	ErrUnauthorized   = New(CodeUnauthorized, "User unauthorized")

	// login //JWT errors
	ErrEmptyToken     = New(CodeUnauthorized, "Token is empty")
	ErrInvalidJWTToken = New(CodeUnauthorized, "Invalid JWT token")
	ErrParseJWTToken   = New(CodeUnauthorized, "Can't parse JWT token")
	ErrExpiredJWTToken = New(CodeUnauthorized, "Expired JWT token")

	// account errors
	ErrIDIsNotValid   = New(CodeUnauthorized, "Id is not valid")
	ErrNotEnoughBalance = New(CodeInsufficientFunds, "Not enough balance")

	// System errors
	ErrInvalidJsonFormat = New(CodeInvalid, "Invalid JSON format")
	ErrCannotConnectDB   = New(CodeInternal, "Cannot connect to database")
	ErrCanUpdateDB      = New(CodeInternal, "Cannot update database")
)
