package validation

import (
	"context"

	"github.com/onexstack/fastgo/pkg/api/apiserver/v1"
)

func ValidateCreateUserRequest(ctx context.Context, rq *v1.CreateUserRequest) error {
	return nil
}

func ValidateUpdateUserRequest(ctx context.Context, rq *v1.UpdateUserRequest) error {
	return nil
}
