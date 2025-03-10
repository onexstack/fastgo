// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/fastgo. The professional
// version of this repository is https://github.com/onexstack/onex.

package conversion

import (
	"github.com/jinzhu/copier"

	"github.com/onexstack/fastgo/internal/apiserver/model"
	apiv1 "github.com/onexstack/fastgo/pkg/api/apiserver/v1"
)

// UserodelToUserV1 将模型层的 User（用户模型对象）转换为 Protobuf 层的 User（v1 用户对象）.
func UserodelToUserV1(userModel *model.User) *apiv1.User {
	var protoUser apiv1.User
	_ = copier.Copy(&protoUser, userModel)
	return &protoUser
}

// UserV1ToUserodel 将 Protobuf 层的 User（v1 用户对象）转换为模型层的 User（用户模型对象）.
func UserV1ToUserodel(protoUser *apiv1.User) *model.User {
	var userModel model.User
	_ = copier.Copy(&userModel, protoUser)
	return &userModel
}
