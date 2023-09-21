package handler

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/Arshia-Izadyar/Jabama-clone/src/api/helper"
	"github.com/gin-gonic/gin"
)



func Create[Tc, Tr any](ctx *gin.Context, fn func(ctx context.Context, req *Tc)(*Tr, error)) {
	req := new(Tc)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithErrorWithValidationError(int(helper.ValidationError), false, err))
		return
	}
	res, err := fn(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithError(nil, int(helper.InternalError), false, err))
		return 
	}
	ctx.JSON(http.StatusCreated, helper.GenerateResponse(res, int(helper.Success), true))
}

func GetById[Tr any](ctx *gin.Context, fn func(ctx context.Context, id int)(*Tr, error)) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if id <= 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithError(nil, int(helper.Error), false, errors.New("wrong value for id")))
		return
	}

	res, err := fn(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithError(nil, int(helper.Error), false, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.GenerateResponse(res, int(helper.Success), true))
}

func Update[Tu, Tr any](ctx *gin.Context, fn func(ctx context.Context, req *Tu, id int)(*Tr, error)) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if id <= 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithError(nil, int(helper.Error), false, errors.New("wrong value for id")))
		return
	}
	req := new(Tu)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithErrorWithValidationError( int(helper.ValidationError), false, err))
		return
	}
	res, err := fn(ctx, req, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithError(nil, int(helper.Error), false, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateResponse(res, int(helper.Success), true))
	
}

func Delete(ctx *gin.Context, fn func(ctx context.Context, id int) error) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if id <= 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithError(nil, int(helper.Error), false, errors.New("wrong value for id")))
		return
	}
	err := fn(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateResponseWithError(nil, int(helper.Error), false, err))
		return
	}
	ctx.JSON(http.StatusNoContent, helper.GenerateResponse(map[string]string{"status":"deleted"}, int(helper.Success), true))
}