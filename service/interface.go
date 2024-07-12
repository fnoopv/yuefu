package service

import (
	"context"

	"github.com/fnoopv/yuefu/dto"
)

type LibraryServiceIface interface {
	Paginate(ctx context.Context, mode string, page, pageSize int) ([]dto.Library, int64, error)
	Create(ctx context.Context, mode, path string) error
	Update(ctx context.Context, id, path string) error
	Delete(ctx context.Context, id string) error
}
