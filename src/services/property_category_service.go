package services

import (
	"context"

	"github.com/Arshia-Izadyar/Jabama-clone/src/api/dto"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/db"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/models"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/logger"
)

type PropertyCategoryService struct {
	base *BaseService[models.PropertyCategory, dto.UpdatePropertyCategoryRequest, dto.CreatePropertyCategoryRequest, dto.PropertyCategoryResponse]
}

func NewPropertyCategoryService(cfg *config.Config) *PropertyCategoryService {
	base := &BaseService[models.PropertyCategory, dto.UpdatePropertyCategoryRequest, dto.CreatePropertyCategoryRequest, dto.PropertyCategoryResponse]{
		DB:  db.GetDB(),
		Log: logger.NewLogger(cfg),
	}
	return &PropertyCategoryService{
		base: base,
	}
}

func (cs *PropertyCategoryService) GetByIdPropertyCategory(ctx context.Context, id int) (*dto.PropertyCategoryResponse, error) {
	return cs.base.GetById(&ctx, id)
}

func (cs *PropertyCategoryService) UpdatePropertyCategory(ctx context.Context, req *dto.UpdatePropertyCategoryRequest, id int) (*dto.PropertyCategoryResponse, error) {
	return cs.base.Update(ctx, req, id)
}

func (cs *PropertyCategoryService) CreatePropertyCategory(ctx context.Context, req *dto.CreatePropertyCategoryRequest) (*dto.PropertyCategoryResponse, error) {
	return cs.base.Create(ctx, req)
}

func (cs *PropertyCategoryService) DeletePropertyCategory(ctx context.Context, id int) error {
	return cs.base.Delete(ctx, id)
}

func (cs *PropertyCategoryService) GetPropertyCategoryByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.PropertyCategoryResponse], error) {
	return cs.base.GetByFilter(ctx, req)
}
