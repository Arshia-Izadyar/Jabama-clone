package services

import (
	"context"

	"github.com/Arshia-Izadyar/Jabama-clone/src/api/dto"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/db"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/models"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/logger"
)

type ResidenceService struct {
	base *BaseService[models.Residence, dto.UpdateResidenceRequest, dto.CreateResidenceRequest, dto.ResidenceResponse]
}

func NewResidenceService(cfg *config.Config) *ResidenceService {
	base := &BaseService[models.Residence, dto.UpdateResidenceRequest, dto.CreateResidenceRequest, dto.ResidenceResponse]{
		DB:  db.GetDB(),
		Log: logger.NewLogger(cfg),
		Preloads: []preload{
			{name: "City"},
			{name: "Province"},
			{name: "RoomType"},
			{name: "ResidenceComment.User"},
			{name: "ResidenceProperties.Property.Category"},
		},
	}
	return &ResidenceService{
		base: base,
	}
}

func (s *ResidenceService) GetByIdResidence(ctx context.Context, id int) (*dto.ResidenceResponse, error) {
	return s.base.GetById(&ctx, id)
}

func (s *ResidenceService) UpdateResidence(ctx context.Context, req *dto.UpdateResidenceRequest, id int) (*dto.ResidenceResponse, error) {
	return s.base.Update(ctx, req, id)
}

func (s *ResidenceService) CreateResidence(ctx context.Context, req *dto.CreateResidenceRequest) (*dto.ResidenceResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *ResidenceService) DeleteResidence(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *ResidenceService) GetResidenceByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.ResidenceResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
