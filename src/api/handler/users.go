package handler

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/Arshia-Izadyar/Jabama-clone/src/api/dto"
	"github.com/Arshia-Izadyar/Jabama-clone/src/api/helper"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/constants"
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

func (uh *UserHandler) RegisterLoginByPhoneNumber(ctx *gin.Context) {
	req := &dto.RegisterLoginByPhone{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithError(nil, -1, false, err))
		return
	}
	res, err := uh.service.RegisterLoginByPhoneNumber(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithError(nil, -1, false, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateResponse(res, 0, true))
}

func (uh *UserHandler) GetOtp(ctx *gin.Context) {
	req := &dto.OtpRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithErrorWithValidationError(-1, false, err))
		return
	}
	err = uh.service.SendOtp(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithError(nil, -1, false, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateResponse(map[string]string{"status": "sent"}, 0, true))

}

func (uh *UserHandler) LoginByUsername(ctx *gin.Context) {
	req := &dto.LoginByUserName{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithErrorWithValidationError(-1, false, err))
		return
	}
	res, err := uh.service.LoginByUserName(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithError(nil, -1, false, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateResponse(res, 0, true))

}

func (uh *UserHandler) RefreshToken(ctx *gin.Context) {
	req := dto.RefreshTokenDTO{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithError(nil, -1, false, err))
		return
	}
	tk, err := uh.service.Token.ValidateRefreshToken(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithError(nil, -1, false, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateResponse(tk, 0, true))
}

func (uh *UserHandler) Logout(ctx *gin.Context) {
	tk := ctx.GetHeader(constants.AuthenticationKey)
	token := strings.Split(tk, " ")
	if len(token) != 2 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithError(nil, -1, false, errors.New("invalid token")))
		return
	}
	err := services.AddToBlacklist(token[1], uh.service.Otp.Cfg.JWT.AccessTokenExpireDuration*time.Minute)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithError(nil, -1, false, err))
		return
	}
	ctx.JSON(http.StatusTeapot, helper.GenerateResponse(map[string]string{"Status": "logged Out"}, 0, true))
}

func (uh *UserHandler) ShowProfile(ctx *gin.Context) {
	res, err := uh.service.ShowUser(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithError(nil, -1, false, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateResponse(res, 0, true))

}
