package repository

import (
	"context"

	"github.com/fnoopv/yuefu/model"
)

type LibrayRepositoryIface interface {
	Paginate(ctx context.Context, mode string, page, pageSize int) ([]model.Library, int64, error)
	Create(ctx context.Context, id, mode, path string) error
	Update(ctx context.Context, id, path string) error
	Delete(ctx context.Context, id string) error
}
