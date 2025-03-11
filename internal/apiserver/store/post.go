// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/fastgo. The professional
// version of this repository is https://github.com/onexstack/onex.

// nolint: dupl
package store

import (
	"context"
	"errors"
	"log/slog"

	"github.com/onexstack/onexstack/pkg/store/where"
	"gorm.io/gorm"

	"github.com/onexstack/fastgo/internal/apiserver/model"
	"github.com/onexstack/fastgo/internal/pkg/errorsx"
)

// PostStore 定义了 post 模块在 store 层所实现的方法.
type PostStore interface {
	Create(ctx context.Context, obj *model.Post) error
	Update(ctx context.Context, obj *model.Post) error
	Delete(ctx context.Context, opts *where.Options) error
	Get(ctx context.Context, opts *where.Options) (*model.Post, error)
	List(ctx context.Context, opts *where.Options) (int64, []*model.Post, error)

	PostExpansion
}

// PostExpansion 定义了帖子操作的附加方法.
type PostExpansion interface{}

// postStore 是 PostStore 接口的实现.
type postStore struct {
	store *datastore
}

// 确保 postStore 实现了 PostStore 接口.
var _ PostStore = (*postStore)(nil)

// newPostStore 创建 postStore 的实例.
func newPostStore(store *datastore) *postStore {
	return &postStore{store}
}

// Create 插入一条帖子记录.
func (s *postStore) Create(ctx context.Context, obj *model.Post) error {
	if err := s.store.DB(ctx).Create(&obj).Error; err != nil {
		slog.Error("Failed to insert post into database", "err", err, "post", obj)
		return errorsx.ErrDBWrite.WithMessage(err.Error())
	}

	return nil
}

// Update 更新帖子数据库记录.
func (s *postStore) Update(ctx context.Context, obj *model.Post) error {
	if err := s.store.DB(ctx).Save(obj).Error; err != nil {
		slog.Error("Failed to update post in database", "err", err, "post", obj)
		return errorsx.ErrDBWrite.WithMessage(err.Error())
	}

	return nil
}

// Delete 根据条件删除帖子记录.
func (s *postStore) Delete(ctx context.Context, opts *where.Options) error {
	err := s.store.DB(ctx, opts).Delete(new(model.Post)).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error("Failed to delete post from database", "err", err, "conditions", opts)
		return errorsx.ErrDBWrite.WithMessage(err.Error())
	}

	return nil
}

// Get 根据条件查询帖子记录.
func (s *postStore) Get(ctx context.Context, opts *where.Options) (*model.Post, error) {
	var obj model.Post
	if err := s.store.DB(ctx, opts).First(&obj).Error; err != nil {
		slog.Error("Failed to retrieve post from database", "err", err, "conditions", opts)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorsx.ErrPostNotFound
		}
		return nil, errorsx.ErrDBRead.WithMessage(err.Error())
	}

	return &obj, nil
}

// List 返回帖子列表和总数.
// nolint: nonamedreturns
func (s *postStore) List(ctx context.Context, opts *where.Options) (count int64, ret []*model.Post, err error) {
	err = s.store.DB(ctx, opts).Order("id desc").Find(&ret).Offset(-1).Limit(-1).Count(&count).Error
	if err != nil {
		slog.Error("Failed to list posts from database", "err", err, "conditions", opts)
		err = errorsx.ErrDBRead.WithMessage(err.Error())
	}
	return
}
