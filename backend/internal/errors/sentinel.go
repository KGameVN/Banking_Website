// sentinel errors (common domain errors)

package errors

// Sentinel errors (reusable, testable)
var (
	// NOTE: account errors
	ErrAccountNotFound = New(CodeNotFound, "Account not found")
	ErrInsufficient    = New(CodeInsufficientFunds, "Insufficient funds")
	ErrInvalidRequest  = New(CodeInvalid, "Invalid request")
	
	// NOTE: login errors
	ErrUserNotFound    = New(CodeNotFound, "User not found")
	ErrUnauthorized   = New(CodeUnauthorized, "User unauthorized")

	// NOTE: login error //JWT errors
	ErrEmptyToken     = New(CodeUnauthorized, "Token is empty")
	ErrInvalidJWTToken = New(CodeUnauthorized, "Invalid JWT token")
	ErrParseJWTToken   = New(CodeUnauthorized, "Can't parse JWT token")
	ErrExpiredJWTToken = New(CodeUnauthorized, "Expired JWT token")

	// NOTE: account errors
	ErrIDIsNotValid   = New(CodeUnauthorized, "Id is not valid")
	ErrNotEnoughBalance = New(CodeInsufficientFunds, "Not enough balance")

	// NOTE: System errors
	ErrInvalidJsonFormat = New(CodeInvalid, "Invalid JSON format")
	ErrCannotConnectDB   = New(CodeInternal, "Cannot connect to database")
	ErrCanUpdateDB      = New(CodeInternal, "Cannot update database")

)
