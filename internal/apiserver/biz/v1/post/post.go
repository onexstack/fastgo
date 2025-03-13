// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/fastgo. The professional
// version of this repository is https://github.com/onexstack/onex.

package post

//go:generate mockgen -destination mock_post.go -package post github.com/onexstack/fastgo/internal/apiserver/biz/v1/post PostBiz

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/onexstack/onexstack/pkg/store/where"

	"github.com/onexstack/fastgo/internal/apiserver/model"
	"github.com/onexstack/fastgo/internal/apiserver/pkg/conversion"
	"github.com/onexstack/fastgo/internal/apiserver/store"
	"github.com/onexstack/fastgo/internal/pkg/contextx"
	apiv1 "github.com/onexstack/fastgo/pkg/api/apiserver/v1"
)

// PostBiz 定义处理帖子请求所需的方法.
type PostBiz interface {
	Create(ctx context.Context, rq *apiv1.CreatePostRequest) (*apiv1.CreatePostResponse, error)
	Update(ctx context.Context, rq *apiv1.UpdatePostRequest) (*apiv1.UpdatePostResponse, error)
	Delete(ctx context.Context, rq *apiv1.DeletePostRequest) (*apiv1.DeletePostResponse, error)
	Get(ctx context.Context, rq *apiv1.GetPostRequest) (*apiv1.GetPostResponse, error)
	List(ctx context.Context, rq *apiv1.ListPostRequest) (*apiv1.ListPostResponse, error)

	PostExpansion
}

// PostExpansion 定义额外的帖子操作方法.
type PostExpansion interface{}

// postBiz 是 PostBiz 接口的实现.
type postBiz struct {
	store store.IStore
}

// 确保 postBiz 实现了 PostBiz 接口.
var _ PostBiz = (*postBiz)(nil)

// New 创建 postBiz 的实例.
func New(store store.IStore) *postBiz {
	return &postBiz{store: store}
}

// Create 实现 PostBiz 接口中的 Create 方法.
func (b *postBiz) Create(ctx context.Context, rq *apiv1.CreatePostRequest) (*apiv1.CreatePostResponse, error) {
	var postM model.Post
	_ = copier.Copy(&postM, rq)
	postM.UserID = contextx.UserID(ctx)

	if err := b.store.Post().Create(ctx, &postM); err != nil {
		return nil, err
	}

	return &apiv1.CreatePostResponse{PostID: postM.PostID}, nil
}

// Update 实现 PostBiz 接口中的 Update 方法.
func (b *postBiz) Update(ctx context.Context, rq *apiv1.UpdatePostRequest) (*apiv1.UpdatePostResponse, error) {
	whr := where.F("userID", contextx.UserID(ctx), "postID", rq.PostID)
	postM, err := b.store.Post().Get(ctx, whr)
	if err != nil {
		return nil, err
	}

	if rq.Title != nil {
		postM.Title = *rq.Title
	}

	if rq.Content != nil {
		postM.Content = *rq.Content
	}

	if err := b.store.Post().Update(ctx, postM); err != nil {
		return nil, err
	}

	return &apiv1.UpdatePostResponse{}, nil
}

// Delete 实现 PostBiz 接口中的 Delete 方法.
func (b *postBiz) Delete(ctx context.Context, rq *apiv1.DeletePostRequest) (*apiv1.DeletePostResponse, error) {
	whr := where.F("userID", contextx.UserID(ctx), "postID", rq.PostIDs)
	if err := b.store.Post().Delete(ctx, whr); err != nil {
		return nil, err
	}

	return &apiv1.DeletePostResponse{}, nil
}

// Get 实现 PostBiz 接口中的 Get 方法.
func (b *postBiz) Get(ctx context.Context, rq *apiv1.GetPostRequest) (*apiv1.GetPostResponse, error) {
	whr := where.F("userID", contextx.UserID(ctx), "postID", rq.PostID)
	postM, err := b.store.Post().Get(ctx, whr)
	if err != nil {
		return nil, err
	}

	return &apiv1.GetPostResponse{Post: conversion.PostodelToPostV1(postM)}, nil
}

// List 实现 PostBiz 接口中的 List 方法.
func (b *postBiz) List(ctx context.Context, rq *apiv1.ListPostRequest) (*apiv1.ListPostResponse, error) {
	whr := where.F("userID", contextx.UserID(ctx)).P(int(rq.Offset), int(rq.Limit))
	if rq.Title != nil {
		whr = whr.Q("title like ?", "%"+*rq.Title+"%")
	}

	count, postList, err := b.store.Post().List(ctx, whr)
	if err != nil {
		return nil, err
	}

	posts := make([]*apiv1.Post, 0, len(postList))
	for _, post := range postList {
		converted := conversion.PostodelToPostV1(post)
		posts = append(posts, converted)
	}

	return &apiv1.ListPostResponse{TotalCount: count, Posts: posts}, nil
}
