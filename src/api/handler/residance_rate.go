package handler

import (
	"github.com/Arshia-Izadyar/Jabama-clone/src/api/dto"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/services"
	"github.com/gin-gonic/gin"
)

type ResidenceRateHandler struct {
	service *services.ResidenceRateService
}

func NewResidenceRateHandler(cfg *config.Config) *ResidenceRateHandler {
	s := services.NewResidenceRateService(cfg)
	return &ResidenceRateHandler{
		service: s,
	}
}

func (ch *ResidenceRateHandler) CreateResidenceRate(ctx *gin.Context) {

	Create[dto.CreateResidenceRateRequest, dto.ResidenceRateResponse](ctx, ch.service.CreateResidenceRate)
}

func (ch *ResidenceRateHandler) GetById(ctx *gin.Context) {

	GetById[dto.ResidenceRateResponse](ctx, ch.service.GetByIdResidenceRate)
}

func (ch *ResidenceRateHandler) UpdateResidenceRate(ctx *gin.Context) {

	Update[dto.UpdateResidenceRateRequest, dto.ResidenceRateResponse](ctx, ch.service.UpdateResidenceRate)

}

func (ch *ResidenceRateHandler) DeleteResidenceRate(ctx *gin.Context) {

	Delete(ctx, ch.service.DeleteResidenceRate)
}
