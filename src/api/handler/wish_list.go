package handler

import (
	"github.com/Arshia-Izadyar/Jabama-clone/src/api/dto"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/services"
	"github.com/gin-gonic/gin"
)

type UserWishListHandler struct {
	service *services.UserWishListService
}

func NewUserWishListHandler(cfg *config.Config) *UserWishListHandler {
	s := services.NewUserWishListService(cfg)
	return &UserWishListHandler{
		service: s,
	}
}

func (ch *UserWishListHandler) CreateUserWishList(ctx *gin.Context) {

	Create[dto.CreateUserWishListRequest, dto.UserWishListResponse](ctx, ch.service.Create)
}

func (ch *UserWishListHandler) GetById(ctx *gin.Context) {

	GetById[dto.UserWishListResponse](ctx, ch.service.GetById)
}

func (ch *UserWishListHandler) DeleteUserWishList(ctx *gin.Context) {

	Delete(ctx, ch.service.Delete)
}
