package handler

import (
	"github.com/Arshia-Izadyar/Jabama-clone/src/api/dto"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/services"
	"github.com/gin-gonic/gin"
)

type RoomTypeHandler struct {
	service *services.RoomTypeService
}

func NewRoomTypeHandler(cfg *config.Config) *RoomTypeHandler {
	s := services.NewRoomTypeService(cfg)
	return &RoomTypeHandler{
		service: s,
	}
}

func (ch *RoomTypeHandler) CreateRoomType(ctx *gin.Context) {

	Create[dto.CreateRoomTypeRequest, dto.RoomTypeResponse](ctx, ch.service.CreateRoomType)
}

func (ch *RoomTypeHandler) GetById(ctx *gin.Context) {

	GetById[dto.RoomTypeResponse](ctx, ch.service.GetByIdRoomType)
}

func (ch *RoomTypeHandler) UpdateRoomType(ctx *gin.Context) {

	Update[dto.UpdateRoomTypeRequest, dto.RoomTypeResponse](ctx, ch.service.UpdateRoomType)

}

func (ch *RoomTypeHandler) DeleteRoomType(ctx *gin.Context) {

	Delete(ctx, ch.service.DeleteRoomType)
}

func (ch *RoomTypeHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter[dto.PaginationInputWithFilter, dto.RoomTypeResponse](ctx, ch.service.GetRoomTypeByFilter)
}
