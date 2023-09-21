package handler

import (
	"github.com/Arshia-Izadyar/Jabama-clone/src/api/dto"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/services"
	"github.com/gin-gonic/gin"
)

type PropertyHandler struct {
	service *services.PropertyService
}

func NewPropertyHandler(cfg *config.Config) *PropertyHandler {
	s := services.NewPropertyService(cfg)
	return &PropertyHandler{
		service: s,
	}
}

func (ch *PropertyHandler) CreateProperty(ctx *gin.Context) {

	Create[dto.CreatePropertyRequest, dto.PropertyResponse](ctx, ch.service.CreateProperty)
}

func (ch *PropertyHandler) GetById(ctx *gin.Context) {

	GetById[dto.PropertyResponse](ctx, ch.service.GetByIdProperty)
}

func (ch *PropertyHandler) UpdateProperty(ctx *gin.Context) {

	Update[dto.UpdatePropertyRequest, dto.PropertyResponse](ctx, ch.service.UpdateProperty)

}

func (ch *PropertyHandler) DeleteProperty(ctx *gin.Context) {

	Delete(ctx, ch.service.DeleteProperty)
}

func (ch *PropertyHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter[dto.PaginationInputWithFilter, dto.PropertyResponse](ctx, ch.service.GetPropertyByFilter)
}
