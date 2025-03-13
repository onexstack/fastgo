package validation

import (
	"github.com/onexstack/fastgo/internal/apiserver/store"
)

// Validator 是验证逻辑的实现结构体.
type Validator struct {
	// 有些复杂的验证逻辑，可能需要直接查询数据库
	// 这里只是一个举例，如果验证时，有其他依赖的客户端/服务/资源等，
	// 都可以一并注入进来
	store store.IStore
}

// NewValidator 创建一个新的 Validator 实例.
func NewValidator(store store.IStore) *Validator {
	return &Validator{store: store}
}
