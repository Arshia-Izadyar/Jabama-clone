package services

import (
	"context"

	"github.com/Arshia-Izadyar/Jabama-clone/src/api/dto"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/db"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/models"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/logger"
)

type ProvinceService struct {
	base *BaseService[models.Province, dto.UpdateProvinceRequest, dto.CreateProvinceRequest, dto.ProvinceResponse]
}

func NewProvinceService(cfg *config.Config) *ProvinceService {
	base := &BaseService[models.Province, dto.UpdateProvinceRequest, dto.CreateProvinceRequest, dto.ProvinceResponse]{
		DB:  db.GetDB(),
		Log: logger.NewLogger(cfg),
	}
	return &ProvinceService{
		base: base,
	}
}

func (cs *ProvinceService) GetByIdProvince(ctx context.Context, id int) (*dto.ProvinceResponse, error) {
	return cs.base.GetById(&ctx, id)
}

func (cs *ProvinceService) UpdateProvince(ctx context.Context, req *dto.UpdateProvinceRequest, id int) (*dto.ProvinceResponse, error) {
	return cs.base.Update(ctx, req, id)
}

func (cs *ProvinceService) CreateProvince(ctx context.Context, req *dto.CreateProvinceRequest) (*dto.ProvinceResponse, error) {
	return cs.base.Create(ctx, req)
}

func (cs *ProvinceService) DeleteProvince(ctx context.Context, id int) error {
	return cs.base.Delete(ctx, id)
}

func (cs *ProvinceService) GetProvinceByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.ProvinceResponse], error) {
	return cs.base.GetByFilter(ctx, req)
}
