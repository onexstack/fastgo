package errorsx

import (
	"errors"
	"fmt"
)

// ErrorX 定义了 fastgo 项目中使用的错误类型，用于描述错误的详细信息.
type ErrorX struct {
	// Code 表示错误的 HTTP 状态码，用于与客户端进行交互时标识错误的类型.
	Code int `json:"code,omitempty"`

	// Reason 表示错误发生的原因，通常为业务错误码，用于精准定位问题.
	Reason string `json:"reason,omitempty"`

	// Message 表示简短的错误信息，通常可直接暴露给用户查看.
	Message string `json:"message,omitempty"`
}

// New 创建一个新的错误.
func New(code int, reason string, format string, args ...any) *ErrorX {
	return &ErrorX{
		Code:    code,
		Reason:  reason,
		Message: fmt.Sprintf(format, args...),
	}
}

// Error 实现 error 接口中的 `Error` 方法.
func (err *ErrorX) Error() string {
	return fmt.Sprintf("error: code = %d reason = %s message = %s", err.Code, err.Reason, err.Message)
}

// WithMessage 设置错误的 Message 字段.
func (err *ErrorX) WithMessage(format string, args ...any) *ErrorX {
	err.Message = fmt.Sprintf(format, args...)
	return err
}

// FromError 尝试将一个通用的 error 转换为自定义的 *ErrorX 类型.
func FromError(err error) *ErrorX {
	// 如果传入的错误是 nil，则直接返回 nil，表示没有错误需要处理.
	if err == nil {
		return nil
	}

	// 检查传入的 error 是否已经是 ErrorX 类型的实例.
	// 如果错误可以通过 errors.As 转换为 *ErrorX 类型，则直接返回该实例.
	if errx := new(ErrorX); errors.As(err, &errx) {
		return errx
	}

	// 默认返回未知错误错误. 该错误代表服务端出错
	return New(ErrInternal.Code, ErrInternal.Reason, err.Error())
}
