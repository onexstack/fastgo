// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/fastgo. The professional
// version of this repository is https://github.com/onexstack/onex.

package store

import (
	"log/slog"
)

// Logger is a logger that implements the Logger interface.
// It uses the log package to log error messages with additional context.
type Logger struct{}

// NewLogger creates and returns a new instance of Logger.
func NewLogger() *Logger {
	return &Logger{}
}

// Error logs an error message with the provided context using the log package.
func (l *Logger) Error(err error, msg string, kvs ...any) {
	slog.Error(msg, append(kvs, "err", err)...)
}
