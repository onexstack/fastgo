// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/fastgo. The professional
// version of this repository is https://github.com/onexstack/onex.

/*
Package contextx 提供了对上下文（context）的扩展功能，允许在 context 中存储和提取用户相关的信息，如用户ID、用户名和访问令牌。

使用后缀 x 表示扩展或变体，使得包名简洁且易于记忆。本包中的函数方便了在上下文中传递和管理用户信息，适用于需要上下文传递数据的场景。

典型用法：
在处理 HTTP 请求的中间件或服务函数中，可以使用这些方法将用户信息存储到上下文中，以便在整个请求生命周期内安全地共享，避免使用全局变量和参数传参。

示例：

	// 创建新的上下文
	ctx := context.Background()

	// 将用户ID和用户名存放到上下文中
	ctx = contextx.WithUserID(ctx, "user-xxxx")
	ctx = contextx.WithUsername(ctx, "sampleUser")

	// 从上下文中提取用户信息
	userID := contextx.UserID(ctx)
	username := contextx.Username(ctx)
*/
package contextx // import "github.com/onexstack/fastgo/internal/pkg/contextx"
