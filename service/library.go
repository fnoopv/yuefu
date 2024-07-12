package service

import (
	"context"

	"github.com/fnoopv/yuefu/dto"
	"github.com/fnoopv/yuefu/repository"
	"github.com/jinzhu/copier"
	"github.com/oklog/ulid/v2"
)

type LibraryServiceImpl struct {
	repository repository.LibrayRepositoryIface
}

func NewLibraryServiceImpl(repository repository.LibrayRepositoryIface) *LibraryServiceImpl {
	return &LibraryServiceImpl{
		repository: repository,
	}
}

func (li *LibraryServiceImpl) Paginate(ctx context.Context, mode string, page int, pageSize int) ([]dto.Library, int64, error) {
	libraryModels, total, err := li.repository.Paginate(ctx, mode, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	var libraries []dto.Library

	if err := copier.Copy(&libraries, &libraryModels); err != nil {
		return nil, 0, err
	}

	return libraries, total, nil
}

func (li *LibraryServiceImpl) Create(ctx context.Context, mode string, path string) error {
	return li.repository.Create(ctx, ulid.Make().String(), mode, path)
}

func (li *LibraryServiceImpl) Update(ctx context.Context, id string, path string) error {
	return li.repository.Update(ctx, id, path)
}

func (li *LibraryServiceImpl) Delete(ctx context.Context, id string) error {
	return li.repository.Delete(ctx, id)
}
