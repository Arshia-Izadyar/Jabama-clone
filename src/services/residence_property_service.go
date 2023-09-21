package services

import (
	"context"

	"github.com/Arshia-Izadyar/Jabama-clone/src/api/dto"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/db"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/models"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/logger"
)

type ResidencePropertyService struct {
	base *BaseService[models.ResidenceProperty, dto.UpdateResidencePropertyRequest, dto.CreateResidencePropertyRequest, dto.ResidencePropertyResponse]
}

func NewResidencePropertyService(cfg *config.Config) *ResidencePropertyService {
	base := &BaseService[models.ResidenceProperty, dto.UpdateResidencePropertyRequest, dto.CreateResidencePropertyRequest, dto.ResidencePropertyResponse]{
		DB:       db.GetDB(),
		Log:      logger.NewLogger(cfg),
		Preloads: []preload{{name: "Property.Category"}},
	}
	return &ResidencePropertyService{
		base: base,
	}
}

func (cs *ResidencePropertyService) GetByIdResidenceProperty(ctx context.Context, id int) (*dto.ResidencePropertyResponse, error) {
	return cs.base.GetById(&ctx, id)
}

func (cs *ResidencePropertyService) UpdateResidenceProperty(ctx context.Context, req *dto.UpdateResidencePropertyRequest, id int) (*dto.ResidencePropertyResponse, error) {
	return cs.base.Update(ctx, req, id)
}

func (cs *ResidencePropertyService) CreateResidenceProperty(ctx context.Context, req *dto.CreateResidencePropertyRequest) (*dto.ResidencePropertyResponse, error) {
	return cs.base.Create(ctx, req)
}

func (cs *ResidencePropertyService) DeleteResidenceProperty(ctx context.Context, id int) error {
	return cs.base.Delete(ctx, id)
}

func (cs *ResidencePropertyService) GetResidencePropertyByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.ResidencePropertyResponse], error) {
	return cs.base.GetByFilter(ctx, req)
}
