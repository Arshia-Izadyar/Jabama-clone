package handler

import (
	"github.com/Arshia-Izadyar/Jabama-clone/src/api/dto"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/services"
	"github.com/gin-gonic/gin"
)

type PropertyCategoryHandler struct {
	service *services.PropertyCategoryService
}

func NewPropertyCategoryHandler(cfg *config.Config) *PropertyCategoryHandler {
	s := services.NewPropertyCategoryService(cfg)
	return &PropertyCategoryHandler{
		service: s,
	}
}

func (ch *PropertyCategoryHandler) CreatePropertyCategory(ctx *gin.Context) {

	Create[dto.CreatePropertyCategoryRequest, dto.PropertyCategoryResponse](ctx, ch.service.CreatePropertyCategory)
}

func (ch *PropertyCategoryHandler) GetById(ctx *gin.Context) {

	GetById[dto.PropertyCategoryResponse](ctx, ch.service.GetByIdPropertyCategory)
}

func (ch *PropertyCategoryHandler) UpdatePropertyCategory(ctx *gin.Context) {

	Update[dto.UpdatePropertyCategoryRequest, dto.PropertyCategoryResponse](ctx, ch.service.UpdatePropertyCategory)

}

func (ch *PropertyCategoryHandler) DeletePropertyCategory(ctx *gin.Context) {

	Delete(ctx, ch.service.DeletePropertyCategory)
}

func (ch *PropertyCategoryHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter[dto.PaginationInputWithFilter, dto.PropertyCategoryResponse](ctx, ch.service.GetPropertyCategoryByFilter)
}
