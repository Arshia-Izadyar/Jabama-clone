package handler

import (
	"net/http"

	"github.com/Arshia-Izadyar/Jabama-clone/src/api/dto"
	"github.com/Arshia-Izadyar/Jabama-clone/src/api/helper"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(cfg *config.Config) *UserHandler {
	s := services.NewUserService(cfg)
	return &UserHandler{
		service: s,
	}
}

func (uh *UserHandler) CreateUserByUsername(ctx *gin.Context) {
	req := &dto.RegisterByUsername{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithError(nil, -1, false, err))
		return
	}
	err = uh.service.RegisterByUsername(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithError(nil, -1, false, err))
		return
	}
	ctx.JSON(http.StatusCreated, helper.GenerateResponse(map[string]string{"Status": "created"}, 0, true))
}
