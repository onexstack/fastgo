package errorsx

import "net/http"

// errorsx 预定义标准的错误.
var (
	// OK 代表请求成功.
	OK = &ErrorX{Code: http.StatusOK, Message: ""}

	// ErrInternal 表示所有未知的服务器端错误.
	ErrInternal = &ErrorX{Code: http.StatusInternalServerError, Reason: "InternalError", Message: "Internal server error."}

	// ErrNotFound 表示资源未找到.
	ErrNotFound = &ErrorX{Code: http.StatusNotFound, Reason: "NotFound", Message: "Resource not found."}

	// ErrDBRead 表示数据库读取失败.
	ErrDBRead = &ErrorX{Code: http.StatusInternalServerError, Reason: "InternalError.DBRead", Message: "Database read failure."}

	// ErrDBWrite 表示数据库写入失败.
	ErrDBWrite = &ErrorX{Code: http.StatusInternalServerError, Reason: "InternalError.DBWrite", Message: "Database write failure."}

	// ErrBind 表示请求体绑定错误.
	ErrBind = &ErrorX{Code: http.StatusBadRequest, Reason: "BindError", Message: "Error occurred while binding the request body to the struct."}

	// ErrInvalidArgument 表示参数验证失败.
	ErrInvalidArgument = &ErrorX{Code: http.StatusBadRequest, Reason: "InvalidArgument", Message: "Argument verification failed."}
)
