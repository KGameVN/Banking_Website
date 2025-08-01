// sentinel errors (common domain errors)

package errors

// Sentinel errors (reusable, testable)
var (
	ErrUserNotFound    = New(CodeNotFound, "user not found")
	ErrAccountNotFound = New(CodeNotFound, "account not found")
	ErrInsufficient    = New(CodeInsufficientFunds, "insufficient funds")
	ErrInvalidRequest  = New(CodeInvalid, "invalid request")
)
