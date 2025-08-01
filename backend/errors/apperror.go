 // AppError type, Wrap, Unwrap, ToAppError, MarshalJSON, Stack

package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"runtime"
	"strings"
)

type AppError struct {
	Code    Code    `json:"code"`
	Message string  `json:"message"`
	Err     error   `json:"-"`      // underlying error (not exposed in JSON)
	trace   []uintptr `json:"-"`    // pc frames
}

// New: create sentinel AppError (no stack, used as sentinel)
func New(code Code, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

// Wrap: wrap underlying error into AppError and capture stack
func Wrap(err error, code Code, message string) *AppError {
	if err == nil {
		return New(code, message)
	}
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
		trace:   callers(),
	}
}

func Wrapf(err error, code Code, format string, args ...interface{}) *AppError {
	return Wrap(err, code, fmt.Sprintf(format, args...))
}

func (e *AppError) Error() string {
	if e == nil {
		return "<nil>"
	}
	if e.Err != nil {
		return fmt.Sprintf("code=%d msg=%s cause=%v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("code=%d msg=%s", e.Code, e.Message)
}

// Unwrap enables errors.Unwrap / errors.Is / errors.As
func (e *AppError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Err
}

// Is allows errors.Is to compare AppError by Code when target is *AppError
func (e *AppError) Is(target error) bool {
	if e == nil || target == nil {
		return false
	}
	var t *AppError
	if errors.As(target, &t) {
		return e.Code == t.Code
	}
	return false
}

// Stack returns human readable stack (if captured)
func (e *AppError) Stack() string {
	if e == nil || len(e.trace) == 0 {
		return ""
	}
	var sb strings.Builder
	frames := runtime.CallersFrames(e.trace)
	for {
		f, more := frames.Next()
		sb.WriteString(fmt.Sprintf("%s\n\t%s:%d\n", f.Function, f.File, f.Line))
		if !more {
			break
		}
	}
	return sb.String()
}

// callers captures stack PCs
func callers() []uintptr {
	const depth = 32
	pcs := make([]uintptr, depth)
	n := runtime.Callers(3, pcs) // skip runtime.Callers, callers, Wrap/Wrapf
	return pcs[:n]
}

// ToAppError converts any error into *AppError (wrap non-AppError as internal)
func ToAppError(err error) *AppError {
	if err == nil {
		return nil
	}
	var ae *AppError
	if errors.As(err, &ae) {
		return ae
	}
	return Wrap(err, CodeInternal, err.Error())
}

// MarshalJSON: only expose safe fields to client (code + message)
func (e *AppError) MarshalJSON() ([]byte, error) {
	type payload struct {
		Code    Code   `json:"code"`
		Message string `json:"message"`
	}
	if e == nil {
		return json.Marshal(payload{Code: CodeUnknown, Message: ""})
	}
	return json.Marshal(payload{Code: e.Code, Message: e.Message})
}
