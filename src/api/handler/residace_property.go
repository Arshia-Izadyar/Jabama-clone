package handler

import (
	"github.com/Arshia-Izadyar/Jabama-clone/src/api/dto"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/services"
	"github.com/gin-gonic/gin"
)

type ResidencePropertyHandler struct {
	service *services.ResidencePropertyService
}

func NewResidencePropertyHandler(cfg *config.Config) *ResidencePropertyHandler {
	s := services.NewResidencePropertyService(cfg)
	return &ResidencePropertyHandler{
		service: s,
	}
}

func (ch *ResidencePropertyHandler) CreateResidenceProperty(ctx *gin.Context) {

	Create[dto.CreateResidencePropertyRequest, dto.ResidencePropertyResponse](ctx, ch.service.CreateResidenceProperty)
}

func (ch *ResidencePropertyHandler) GetById(ctx *gin.Context) {

	GetById[dto.ResidencePropertyResponse](ctx, ch.service.GetByIdResidenceProperty)
}

func (ch *ResidencePropertyHandler) UpdateResidenceProperty(ctx *gin.Context) {
	Update[dto.UpdateResidencePropertyRequest, dto.ResidencePropertyResponse](ctx, ch.service.UpdateResidenceProperty)

}

func (ch *ResidencePropertyHandler) DeleteResidenceProperty(ctx *gin.Context) {
	Delete(ctx, ch.service.DeleteResidenceProperty)
}

func (ch *ResidencePropertyHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter[dto.PaginationInputWithFilter, dto.ResidencePropertyResponse](ctx, ch.service.GetResidencePropertyByFilter)
}
