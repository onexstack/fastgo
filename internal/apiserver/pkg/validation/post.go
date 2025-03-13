package validation

import (
	"context"

	"github.com/onexstack/fastgo/pkg/api/apiserver/v1"
)

func (v *Validator) ValidateCreatePostRequest(ctx context.Context, rq *v1.CreatePostRequest) error {
	return nil
}

func (v *Validator) ValidateUpdatePostRequest(ctx context.Context, rq *v1.UpdatePostRequest) error {
	return nil
}
