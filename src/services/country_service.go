package services

import (
	"context"

	"github.com/Arshia-Izadyar/Jabama-clone/src/api/dto"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/db"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/models"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/logger"
)

type CountryService struct {
	base *BaseService[models.City, dto.UpdateCityRequest, dto.CreateCityRequest, dto.CityResponse]
}

func NewCountryService(cfg *config.Config) *CountryService {
	base := &BaseService[models.City, dto.UpdateCityRequest, dto.CreateCityRequest, dto.CityResponse]{
		DB:       db.GetDB(),
		Log:      logger.NewLogger(cfg),
		Preloads: []preload{{name: "Provinces"}},
	}
	return &CountryService{
		base: base,
	}
}

func (cs *CountryService) GetByIdCity(ctx context.Context, id int) (*dto.CityResponse, error) {
	return cs.base.GetById(&ctx, id)
}

func (cs *CountryService) UpdateCity(ctx context.Context, req *dto.UpdateCityRequest, id int) (*dto.CityResponse, error) {
	return cs.base.Update(ctx, req, id)
}

func (cs *CountryService) CreateCity(ctx context.Context, req *dto.CreateCityRequest) (*dto.CityResponse, error) {
	return cs.base.Create(ctx, req)
}

func (cs *CountryService) DeleteCity(ctx context.Context, id int) error {
	return cs.base.Delete(ctx, id)
}

func (cs *CountryService) GetCityByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.CityResponse], error) {
	return cs.base.GetByFilter(ctx, req)
}
