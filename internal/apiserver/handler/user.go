// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/fastgo. The professional
// version of this repository is https://github.com/onexstack/onex.

package handler

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/onexstack/fastgo/internal/apiserver/pkg/validation"
	"github.com/onexstack/fastgo/internal/pkg/core"
	"github.com/onexstack/fastgo/internal/pkg/errorsx"
	"github.com/onexstack/fastgo/pkg/api/apiserver/v1"
)

// CreateUser 创建新用户.
func (h *Handler) CreateUser(c *gin.Context) {
	slog.Info("Create user function called")

	var rq v1.CreateUserRequest
	if err := c.ShouldBindJSON(&rq); err != nil {
		core.WriteResponse(c, errorsx.ErrBind, nil)
		return
	}

	if err := validation.ValidateCreateUserRequest(c.Request.Context(), &rq); err != nil {
		core.WriteResponse(c, errorsx.ErrInvalidArgument.WithMessage(err.Error()), nil)
		return
	}

	resp, err := h.biz.UserV1().Create(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, resp)
}

// UpdateUser 更新用户信息.
func (h *Handler) UpdateUser(c *gin.Context) {
	slog.Info("Update user function called")

	var rq v1.UpdateUserRequest
	if err := c.ShouldBindJSON(&rq); err != nil {
		core.WriteResponse(c, errorsx.ErrBind, nil)
		return
	}

	if err := validation.ValidateUpdateUserRequest(c.Request.Context(), &rq); err != nil {
		core.WriteResponse(c, errorsx.ErrInvalidArgument.WithMessage(err.Error()), nil)
		return
	}

	resp, err := h.biz.UserV1().Update(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, resp)
}

// DeleteUser 删除用户.
func (h *Handler) DeleteUser(c *gin.Context) {
	slog.Info("Delete user function called")

	var rq v1.DeleteUserRequest
	if err := c.ShouldBindUri(&rq); err != nil {
		core.WriteResponse(c, errorsx.ErrBind, nil)
		return
	}

	// 小作业：请你自行补全校验代码

	resp, err := h.biz.UserV1().Delete(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, resp)
}

// GetUser 获取用户信息.
func (h *Handler) GetUser(c *gin.Context) {
	slog.Info("Get user function called")

	var rq v1.GetUserRequest
	if err := c.ShouldBindUri(&rq); err != nil {
		core.WriteResponse(c, errorsx.ErrBind, nil)
		return
	}

	// 小作业：请你自行补全校验代码

	resp, err := h.biz.UserV1().Get(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, resp)
}

// ListUser 列出用户信息.
func (h *Handler) ListUser(c *gin.Context) {
	slog.Info("List user function called")

	var rq v1.ListUserRequest
	if err := c.ShouldBindQuery(&rq); err != nil {
		core.WriteResponse(c, errorsx.ErrBind, nil)
		return
	}

	// 小作业：请你自行补全校验代码

	resp, err := h.biz.UserV1().List(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, resp)
}
