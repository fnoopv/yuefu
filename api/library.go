package api

import (
	"errors"

	"github.com/fnoopv/yuefu/dto"
	"github.com/fnoopv/yuefu/service"
	"github.com/gin-gonic/gin"
)

type LibraryAPI struct {
	service service.LibraryServiceIface
}

func NewLibraryAPI(service service.LibraryServiceIface) *LibraryAPI {
	return &LibraryAPI{
		service: service,
	}
}

func (li *LibraryAPI) List(ctx *gin.Context) {
	var req dto.LibraryList

	if err := ctx.ShouldBindQuery(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	libraries, total, err := li.service.Paginate(ctx.Request.Context(), req.Mode, req.Page, req.PageSize)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Set("data", dto.PaginationDataEntry{
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
		List:     libraries,
	})
}

func (li *LibraryAPI) Create(ctx *gin.Context) {
	var req dto.LibraryCreate

	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := li.service.Create(ctx.Request.Context(), req.Mode, req.Path); err != nil {
		_ = ctx.Error(err)
		return
	}
}

func (li *LibraryAPI) Update(ctx *gin.Context) {
	var req dto.LibraryUpdate

	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := li.service.Update(ctx.Request.Context(), req.ID, req.Path); err != nil {
		_ = ctx.Error(err)
		return
	}
}

func (li *LibraryAPI) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		_ = ctx.Error(errors.New("id is not exist"))
		return
	}

	if err := li.service.Delete(ctx.Request.Context(), id); err != nil {
		_ = ctx.Error(err)
		return
	}
}
