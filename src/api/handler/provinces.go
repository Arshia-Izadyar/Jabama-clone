package handler

import (
	"github.com/Arshia-Izadyar/Jabama-clone/src/api/dto"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/services"
	"github.com/gin-gonic/gin"
)


type ProvinceHandler struct {
	service *services.ProvinceService
}

func NewProvinceHandler(cfg *config.Config) *ProvinceHandler {
	s := services.NewProvinceService(cfg)
	return &ProvinceHandler{
		service: s,
	}
}

func (ch *ProvinceHandler) CreateProvince(ctx *gin.Context) {
	
	Create[dto.CreateProvinceRequest, dto.ProvinceResponse](ctx, ch.service.CreateProvince)
}

func (ch *ProvinceHandler) GetById(ctx *gin.Context) {
	
	GetById[dto.ProvinceResponse](ctx, ch.service.GetByIdProvince)
}

func (ch *ProvinceHandler) UpdateProvince(ctx *gin.Context) {
	
	Update[dto.UpdateProvinceRequest, dto.ProvinceResponse](ctx, ch.service.UpdateProvince)

}

func (ch *ProvinceHandler) DeleteProvince(ctx *gin.Context) {

	Delete(ctx, ch.service.DeleteProvince)
}