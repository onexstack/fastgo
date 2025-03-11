// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/fastgo. The professional
// version of this repository is https://github.com/onexstack/onex.

package user

//go:generate mockgen -destination mock_user.go -package user github.com/onexstack/fastgo/internal/apiserver/biz/v1/user UserBiz

import (
	"context"
	"log/slog"
	"sync"

	"github.com/jinzhu/copier"
	"github.com/onexstack/onexstack/pkg/store/where"
	"golang.org/x/sync/errgroup"

	"github.com/onexstack/fastgo/internal/apiserver/model"
	"github.com/onexstack/fastgo/internal/apiserver/pkg/conversion"
	"github.com/onexstack/fastgo/internal/apiserver/store"
	"github.com/onexstack/fastgo/internal/pkg/contextx"
	"github.com/onexstack/fastgo/internal/pkg/errorsx"
	"github.com/onexstack/fastgo/internal/pkg/known"
	apiv1 "github.com/onexstack/fastgo/pkg/api/apiserver/v1"
	"github.com/onexstack/fastgo/pkg/auth"
	"github.com/onexstack/fastgo/pkg/token"
)

// UserBiz 定义处理用户请求所需的方法.
type UserBiz interface {
	Create(ctx context.Context, rq *apiv1.CreateUserRequest) (*apiv1.CreateUserResponse, error)
	Update(ctx context.Context, rq *apiv1.UpdateUserRequest) (*apiv1.UpdateUserResponse, error)
	Delete(ctx context.Context, rq *apiv1.DeleteUserRequest) (*apiv1.DeleteUserResponse, error)
	Get(ctx context.Context, rq *apiv1.GetUserRequest) (*apiv1.GetUserResponse, error)
	List(ctx context.Context, rq *apiv1.ListUserRequest) (*apiv1.ListUserResponse, error)

	UserExpansion
}

// UserExpansion 定义用户操作的扩展方法.
type UserExpansion interface {
	Login(ctx context.Context, rq *apiv1.LoginRequest) (*apiv1.LoginResponse, error)
	RefreshToken(ctx context.Context, rq *apiv1.RefreshTokenRequest) (*apiv1.RefreshTokenResponse, error)
	ChangePassword(ctx context.Context, rq *apiv1.ChangePasswordRequest) (*apiv1.ChangePasswordResponse, error)
}

// userBiz 是 UserBiz 接口的实现.
type userBiz struct {
	store store.IStore
}

// 确保 userBiz 实现了 UserBiz 接口.
var _ UserBiz = (*userBiz)(nil)

func New(store store.IStore) *userBiz {
	return &userBiz{store: store}
}

// Login 实现 UserBiz 接口中的 Login 方法.
func (b *userBiz) Login(ctx context.Context, rq *apiv1.LoginRequest) (*apiv1.LoginResponse, error) {
	// 获取登录用户的所有信息
	whr := where.F("username", rq.Username)
	userM, err := b.store.User().Get(ctx, whr)
	if err != nil {
		return nil, errorsx.ErrUserNotFound
	}

	// 对比传入的明文密码和数据库中已加密过的密码是否匹配
	if err := auth.Compare(userM.Password, rq.Password); err != nil {
		slog.ErrorContext(ctx, "Failed to compare password", "err", err)
		return nil, errorsx.ErrPasswordInvalid
	}

	// 如果匹配成功，说明登录成功，签发 token 并返回
	tokenStr, expireAt, err := token.Sign(userM.UserID)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to sign token", "err", err)
		return nil, errorsx.ErrSignToken
	}

	return &apiv1.LoginResponse{Token: tokenStr, ExpireAt: expireAt}, nil
}

// RefreshToken 用于刷新用户的身份验证令牌.
// 当用户的令牌即将过期时，可以调用此方法生成一个新的令牌.
func (b *userBiz) RefreshToken(ctx context.Context, rq *apiv1.RefreshTokenRequest) (*apiv1.RefreshTokenResponse, error) {
	// 如果匹配成功，说明登录成功，签发 token 并返回
	tokenStr, expireAt, err := token.Sign(contextx.UserID(ctx))
	if err != nil {
		return nil, errorsx.ErrSignToken.WithMessage(err.Error())
	}
	return &apiv1.RefreshTokenResponse{Token: tokenStr, ExpireAt: expireAt}, nil
}

// ChangePassword 实现 UserBiz 接口中的 ChangePassword 方法.
func (b *userBiz) ChangePassword(ctx context.Context, rq *apiv1.ChangePasswordRequest) (*apiv1.ChangePasswordResponse, error) {
	userM, err := b.store.User().Get(ctx, where.F("userID", contextx.UserID(ctx)))
	if err != nil {
		return nil, err
	}

	if err := auth.Compare(userM.Password, rq.OldPassword); err != nil {
		slog.ErrorContext(ctx, "Failed to compare password", "err", err)
		return nil, errorsx.ErrPasswordInvalid
	}

	userM.Password, _ = auth.Encrypt(rq.NewPassword)
	if err := b.store.User().Update(ctx, userM); err != nil {
		return nil, err
	}

	return &apiv1.ChangePasswordResponse{}, nil
}

// Create 实现 UserBiz 接口中的 Create 方法.
func (b *userBiz) Create(ctx context.Context, rq *apiv1.CreateUserRequest) (*apiv1.CreateUserResponse, error) {
	var userM model.User
	_ = copier.Copy(&userM, rq)

	if err := b.store.User().Create(ctx, &userM); err != nil {
		return nil, err
	}

	return &apiv1.CreateUserResponse{UserID: userM.UserID}, nil
}

// Update 实现 UserBiz 接口中的 Update 方法.
func (b *userBiz) Update(ctx context.Context, rq *apiv1.UpdateUserRequest) (*apiv1.UpdateUserResponse, error) {
	userM, err := b.store.User().Get(ctx, where.F("userID", contextx.UserID(ctx)))
	if err != nil {
		return nil, err
	}

	if rq.Username != nil {
		userM.Username = *rq.Username
	}
	if rq.Email != nil {
		userM.Email = *rq.Email
	}
	if rq.Nickname != nil {
		userM.Nickname = *rq.Nickname
	}
	if rq.Phone != nil {
		userM.Phone = *rq.Phone
	}

	if err := b.store.User().Update(ctx, userM); err != nil {
		return nil, err
	}

	return &apiv1.UpdateUserResponse{}, nil
}

// Delete 实现 UserBiz 接口中的 Delete 方法.
func (b *userBiz) Delete(ctx context.Context, rq *apiv1.DeleteUserRequest) (*apiv1.DeleteUserResponse, error) {
	if err := b.store.User().Delete(ctx, where.F("userID", contextx.UserID(ctx))); err != nil {
		return nil, err
	}

	return &apiv1.DeleteUserResponse{}, nil
}

// Get 实现 UserBiz 接口中的 Get 方法.
func (b *userBiz) Get(ctx context.Context, rq *apiv1.GetUserRequest) (*apiv1.GetUserResponse, error) {
	userM, err := b.store.User().Get(ctx, where.F("userID", contextx.UserID(ctx)))
	if err != nil {
		return nil, err
	}

	return &apiv1.GetUserResponse{User: conversion.UserodelToUserV1(userM)}, nil
}

// List 实现 UserBiz 接口中的 List 方法.
func (b *userBiz) List(ctx context.Context, rq *apiv1.ListUserRequest) (*apiv1.ListUserResponse, error) {
	whr := where.P(int(rq.Offset), int(rq.Limit))
	count, userList, err := b.store.User().List(ctx, whr)
	if err != nil {
		return nil, err
	}

	var m sync.Map
	eg, ctx := errgroup.WithContext(ctx)

	// 设置最大并发数量为常量 MaxConcurrency
	eg.SetLimit(known.MaxErrGroupConcurrency)

	// 使用 goroutine 提高接口性能
	for _, user := range userList {
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				count, _, err := b.store.Post().List(ctx, where.F("userID", contextx.UserID(ctx)))
				if err != nil {
					return err
				}

				converted := conversion.UserodelToUserV1(user)
				converted.PostCount = count
				m.Store(user.ID, converted)

				return nil
			}
		})
	}

	if err := eg.Wait(); err != nil {
		slog.ErrorContext(ctx, "Failed to wait all function calls returned", "err", err)
		return nil, err
	}

	users := make([]*apiv1.User, 0, len(userList))
	for _, item := range userList {
		user, _ := m.Load(item.ID)
		users = append(users, user.(*apiv1.User))
	}

	slog.DebugContext(ctx, "Get users from backend storage", "count", len(users))

	return &apiv1.ListUserResponse{TotalCount: count, Users: users}, nil
}
