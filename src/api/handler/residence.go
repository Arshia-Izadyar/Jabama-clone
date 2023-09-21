package handler

import (
	"github.com/Arshia-Izadyar/Jabama-clone/src/api/dto"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/services"
	"github.com/gin-gonic/gin"
)

type ResidenceHandler struct {
	service *services.ResidenceService
}

func NewResidenceHandler(cfg *config.Config) *ResidenceHandler {
	s := services.NewResidenceService(cfg)
	return &ResidenceHandler{
		service: s,
	}
}

func (ch *ResidenceHandler) CreateResidence(ctx *gin.Context) {

	Create[dto.CreateResidenceRequest, dto.ResidenceResponse](ctx, ch.service.CreateResidence)
}

func (ch *ResidenceHandler) GetById(ctx *gin.Context) {

	GetById[dto.ResidenceResponse](ctx, ch.service.GetByIdResidence)
}

func (ch *ResidenceHandler) UpdateResidence(ctx *gin.Context) {

	Update[dto.UpdateResidenceRequest, dto.ResidenceResponse](ctx, ch.service.UpdateResidence)

}

func (ch *ResidenceHandler) DeleteResidence(ctx *gin.Context) {

	Delete(ctx, ch.service.DeleteResidence)
}

func (ch *ResidenceHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter[dto.PaginationInputWithFilter, dto.ResidenceResponse](ctx, ch.service.GetResidenceByFilter)
}
