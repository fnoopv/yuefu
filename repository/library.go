package repository

import (
	"context"

	"github.com/fnoopv/yuefu/model"
	"gorm.io/gorm"
)

type LibraryRepositoryXormImpl struct {
	db *gorm.DB
}

func NewLibraryRepositoryXormImpl(db *gorm.DB) *LibraryRepositoryXormImpl {
	return &LibraryRepositoryXormImpl{
		db: db,
	}
}

func (li *LibraryRepositoryXormImpl) Paginate(ctx context.Context, mode string, page int, pageSize int) ([]model.Library, int64, error) {
	offset := (page - 1) * pageSize

	var libraries []model.Library
	var total int64

	db := li.db.WithContext(ctx).Model(&model.Library{}).Where("mode = ?", mode)

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := db.Offset(offset).Limit(pageSize).Find(&libraries).Error; err != nil {
		return nil, 0, err
	}

	return libraries, total, nil
}

func (li *LibraryRepositoryXormImpl) Create(ctx context.Context, id string, mode string, path string) error {
	library := model.Library{
		ID:   id,
		Mode: mode,
		Path: path,
	}

	return li.db.WithContext(ctx).Create(&library).Error
}

func (li *LibraryRepositoryXormImpl) Update(ctx context.Context, id string, path string) error {
	library := model.Library{
		ID: id,
	}

	return li.db.WithContext(ctx).Model(&library).Update("path", path).Error
}

func (li *LibraryRepositoryXormImpl) Delete(ctx context.Context, id string) error {
	return li.db.WithContext(ctx).Delete(&model.Library{ID: id}).Error
}
