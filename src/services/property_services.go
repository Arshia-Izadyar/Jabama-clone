package services

import (
	"context"

	"github.com/Arshia-Izadyar/Jabama-clone/src/api/dto"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/db"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/models"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/logger"
)

type PropertyService struct {
	base *BaseService[models.Property, dto.UpdatePropertyRequest, dto.CreatePropertyRequest, dto.PropertyResponse]
}

func NewPropertyService(cfg *config.Config) *PropertyService {
	base := &BaseService[models.Property, dto.UpdatePropertyRequest, dto.CreatePropertyRequest, dto.PropertyResponse]{
		DB:       db.GetDB(),
		Log:      logger.NewLogger(cfg),
		Preloads: []preload{{name: "Category"}},
	}
	return &PropertyService{
		base: base,
	}
}

func (cs *PropertyService) GetByIdProperty(ctx context.Context, id int) (*dto.PropertyResponse, error) {
	return cs.base.GetById(&ctx, id)
}

func (cs *PropertyService) UpdateProperty(ctx context.Context, req *dto.UpdatePropertyRequest, id int) (*dto.PropertyResponse, error) {
	return cs.base.Update(ctx, req, id)
}

func (cs *PropertyService) CreateProperty(ctx context.Context, req *dto.CreatePropertyRequest) (*dto.PropertyResponse, error) {
	return cs.base.Create(ctx, req)
}

func (cs *PropertyService) DeleteProperty(ctx context.Context, id int) error {
	return cs.base.Delete(ctx, id)
}

func (cs *PropertyService) GetPropertyByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.PropertyResponse], error) {
	return cs.base.GetByFilter(ctx, req)
}
