// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/fastgo. The professional
// version of this repository is https://github.com/onexstack/onex.

package rid

import (
	"github.com/onexstack/onexstack/pkg/id"
)

const defaultABC = "abcdefghijklmnopqrstuvwxyz1234567890"

type ResourceID string

const (
	// UserID 定义用户资源标识符.
	UserID ResourceID = "user"
	// PostID 定义博文资源标识符.
	PostID ResourceID = "post"
)

// String 将资源标识符转换为字符串.
func (rid ResourceID) String() string {
	return string(rid)
}

// New 创建带前缀的唯一标识符.
func (rid ResourceID) New(counter uint64) string {
	// 使用自定义选项生成唯一标识符
	uniqueStr := id.NewCode(
		counter,
		id.WithCodeChars([]rune(defaultABC)),
		id.WithCodeL(6),
		id.WithCodeSalt(Salt()),
	)
	return rid.String() + "-" + uniqueStr
}
