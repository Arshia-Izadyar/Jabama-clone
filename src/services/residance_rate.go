package services

import (
	"context"
	"errors"

	"github.com/Arshia-Izadyar/Jabama-clone/src/api/dto"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/constants"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/db"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/models"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/logger"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/service_errors"
	"gorm.io/gorm"
)

type ResidenceRateService struct {
	base *BaseService[models.ResidenceRate, dto.UpdateResidenceRateRequest, dto.CreateResidenceRateRequest, dto.ResidenceRateResponse]
}

func NewResidenceRateService(cfg *config.Config) *ResidenceRateService {
	base := &BaseService[models.ResidenceRate, dto.UpdateResidenceRateRequest, dto.CreateResidenceRateRequest, dto.ResidenceRateResponse]{
		DB:  db.GetDB(),
		Log: logger.NewLogger(cfg),
		// Preloads: []preload{{name: "Property.Category"}},
	}
	return &ResidenceRateService{
		base: base,
	}
}

func (cs *ResidenceRateService) GetByIdResidenceRate(ctx context.Context, id int) (*dto.ResidenceRateResponse, error) {
	return cs.base.GetById(&ctx, id)
}

func (cs *ResidenceRateService) UpdateResidenceRate(ctx context.Context, req *dto.UpdateResidenceRateRequest, id int) (*dto.ResidenceRateResponse, error) {
	return cs.base.Update(ctx, req, id)
}

func (cs *ResidenceRateService) CreateResidenceRate(ctx context.Context, req *dto.CreateResidenceRateRequest) (*dto.ResidenceRateResponse, error) {
	id := int(ctx.Value(constants.UserIdKey).(float64))
	model := &models.ResidenceRate{}
	err := cs.base.DB.Model(&models.ResidenceRate{}).Where("residence_id = ?", req.ResidenceId).Where("user_id = ?", id).First(&model).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		req.UserId = id
		return cs.base.Create(ctx, req)
	} else {
		return nil, &service_errors.ServiceError{EndUserMessage: "users can only rate once per residence", Err: errors.New("users cant rate more than onr time")}
	}
}

func (cs *ResidenceRateService) DeleteResidenceRate(ctx context.Context, id int) error {
	return cs.base.Delete(ctx, id)
}
