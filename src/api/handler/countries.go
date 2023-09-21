package handler

import (
	"github.com/Arshia-Izadyar/Jabama-clone/src/api/dto"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/services"
	"github.com/gin-gonic/gin"
)

type CityHandler struct {
	service *services.CountryService
}

func NewCityHandler(cfg *config.Config) *CityHandler {
	s := services.NewCountryService(cfg)
	return &CityHandler{
		service: s,
	}
}

func (ch *CityHandler) CreateCity(ctx *gin.Context) {
	// req := &dto.CreateCityRequest{}
	// err := ctx.ShouldBindJSON(&req)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithErrorWithValidationError(-1, false, err))
	// 	return
	// }
	// res , err := ch.service.CreateCity(ctx, req)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithError(nil, -1, false, err))
	// 	return
	// }

	// ctx.JSON(http.StatusCreated, helper.GenerateResponse(res, 0, true))
	Create[dto.CreateCityRequest, dto.CityResponse](ctx, ch.service.CreateCity)
}

func (ch *CityHandler) GetById(ctx *gin.Context) {
	// id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	// if id <= 0 {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithError(nil, -1, false, errors.New("wrong value for id")))
	// 	return
	// }

	// res, err := ch.service.GetByIdCity(ctx, id)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithError(nil, -1, false, err))
	// 	return
	// }
	// ctx.JSON(http.StatusOK, helper.GenerateResponse(res, 0, true))
	GetById[dto.CityResponse](ctx, ch.service.GetByIdCity)
}

func (ch *CityHandler) UpdateCity(ctx *gin.Context) {
	// id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	// if id <= 0 {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithError(nil, -1, false, errors.New("wrong value for id")))
	// 	return
	// }
	// req := dto.UpdateCityRequest{}
	// err := ctx.ShouldBindJSON(&req)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithErrorWithValidationError(-1, false, err))
	// 	return
	// }

	// res, err := ch.service.UpdateCity(ctx, &req, id)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithError(nil, -1, false, err))
	// 	return
	// }
	// ctx.JSON(http.StatusOK, helper.GenerateResponse(res, 0, true))
	Update[dto.UpdateCityRequest, dto.CityResponse](ctx, ch.service.UpdateCity)

}

func (ch *CityHandler) DeleteCity(ctx *gin.Context) {
	// id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	// if id <= 0 {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithError(nil, -1, false, errors.New("wrong value for id")))
	// 	return
	// }

	// err := ch.service.DeleteCity(ctx, id)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithError(nil, -1, false, err))
	// 	return
	// }
	// ctx.JSON(http.StatusNoContent, helper.GenerateResponse(map[string]string{"status":"Deleted"}, 0, true))
	Delete(ctx, ch.service.DeleteCity)
}

func (ch *CityHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter[dto.PaginationInputWithFilter, dto.CityResponse](ctx, ch.service.GetCityByFilter)
}
