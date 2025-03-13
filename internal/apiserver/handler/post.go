// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/fastgo. The professional
// version of this repository is https://github.com/onexstack/onex.

package handler

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/onexstack/fastgo/internal/pkg/core"
	"github.com/onexstack/fastgo/internal/pkg/errorsx"
	"github.com/onexstack/fastgo/pkg/api/apiserver/v1"
)

// CreatePost 创建新博客.
func (h *Handler) CreatePost(c *gin.Context) {
	slog.Info("Create post function called")

	var rq v1.CreatePostRequest
	if err := c.ShouldBindJSON(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)
		return
	}

	if err := h.val.ValidateCreatePostRequest(c.Request.Context(), &rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrInvalidArgument.WithMessage(err.Error()))
		return
	}

	resp, err := h.biz.PostV1().Create(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

// UpdatePost 更新博客信息.
func (h *Handler) UpdatePost(c *gin.Context) {
	slog.Info("Update post function called")

	var rq v1.UpdatePostRequest
	if err := c.ShouldBindJSON(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)
		return
	}
	rq.PostID = c.Param("postID")

	if err := h.val.ValidateUpdatePostRequest(c.Request.Context(), &rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrInvalidArgument.WithMessage(err.Error()))
		return
	}

	resp, err := h.biz.PostV1().Update(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

// DeletePost 删除博客.
func (h *Handler) DeletePost(c *gin.Context) {
	slog.Info("Delete post function called")

	var rq v1.DeletePostRequest
	if err := c.ShouldBindJSON(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)
		return
	}

	// 小作业：请你自行补全校验代码

	resp, err := h.biz.PostV1().Delete(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

// GetPost 获取博客信息.
func (h *Handler) GetPost(c *gin.Context) {
	slog.Info("Get post function called")

	var rq v1.GetPostRequest
	if err := c.ShouldBindUri(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)
		return
	}

	// 小作业：请你自行补全校验代码

	resp, err := h.biz.PostV1().Get(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

// ListPost 列出博客信息.
func (h *Handler) ListPost(c *gin.Context) {
	slog.Info("List post function called")

	var rq v1.ListPostRequest
	if err := c.ShouldBindQuery(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)
		return
	}

	// 小作业：请你自行补全校验代码

	resp, err := h.biz.PostV1().List(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}
