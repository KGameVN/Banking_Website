// mapping AppError -> HTTP response payload (and logs)

package errors

// HTTPStatusAndPayload returns (status, payload) safe for returning to client
func HTTPStatusAndPayload(err error) (int, interface{}) {
	appError := ToAppError(err)
	if appError == nil {
		// unexpected nil -> internal server error
		return CodeInternal , map[string]interface{}{
			"code":    int(CodeInternal),
			"message": "internal server error",
		}
	}
	status := appError.Code.HTTPStatus()
	msg := appError.Message
	// Do not leak internal message to client for CodeInternal
	if appError.Code == CodeInternal {
		msg = "internal server error"
	}
	// TODO: Log the error with stack trace here
	//
	//
	return status, map[string]interface{}{
		"code":    int(appError.Code),
		"message": msg,
	}
}
