package error

import "fmt"

// DefaultError đại diện cho lỗi cơ bản với mã và thông điệp
type DefaultError struct {
	Code    int
	Message string
}

// SysError là lỗi tùy chỉnh kết hợp DefaultError và lỗi gốc
type SysError struct {
	Info DefaultError
	Err  error
}

// NewSysError khởi tạo một SysError mới
func NewSysError(mess string, code int, err error) *SysError {
	return &SysError{
		Info: DefaultError{
			Code:    code,
			Message: mess,
		},
		Err: err,
	}
}

// Error triển khai phương thức của giao diện error, trả về thông điệp lỗi
func (e *SysError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s, Original Error: %s", e.Info.Code, e.Info.Message, e.Err.Error())
}
