package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Arshia-Izadyar/Jabama-clone/src/api/helper"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/constants"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/service_errors"
	"github.com/Arshia-Izadyar/Jabama-clone/src/services"
	"github.com/gin-gonic/gin"
)


func Authentication(cfg *config.Config) gin.HandlerFunc {
	var tokenService = services.NewTokenService(cfg)
	var err error
	return func(ctx *gin.Context) {
		claimMap := map[string]interface{}{}
		key := ctx.GetHeader(constants.AuthenticationKey)
		if key == ""{
			err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenNotPresent}
		} else {
			token := strings.Split(key, " ")[1]
			claimMap, err = tokenService.GetClaims(token)
			if err != nil {
				e, isServiceError := err.(*service_errors.ServiceError)
				if isServiceError && e.EndUserMessage == service_errors.TokenExpired {
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.GenerateResponseWithError(nil, -1, false, err))
					return
				}else {
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.GenerateResponseWithError(nil, -1, false, err))
					return
			} 
			}

		}
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.GenerateResponseWithError(nil, -1, false, err))
			return
		}
		ctx.Set(constants.UserIdKey, claimMap[constants.UserIdKey])
		ctx.Set(constants.PhoneKey, claimMap[constants.PhoneKey])
		ctx.Set(constants.ExpKey, claimMap[constants.ExpKey])
		ctx.Set(constants.UsernameKey, claimMap[constants.UsernameKey])
		ctx.Set(constants.RolesKey, claimMap[constants.RolesKey])
		
		ctx.Next()
	}
}

func Authorization(validRole []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(ctx.Keys) == 0 {
			ctx.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateResponseWithError(nil, -1, false, errors.New("no token provided")))
			return
		}
		rolesV, ok := ctx.Keys[constants.RolesKey]
		if !ok {
			if len(ctx.Keys) == 0 {
				ctx.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateResponseWithError(nil, -1, false, errors.New("dont have required roles")))
				return
			}
		}
		roles := rolesV.([]interface{})
		val := map[string]struct{}{}
		for _, r := range roles {
			val[r.(string)] = struct{}{}
		} 
		for _, item := range validRole {
			if _, ok := val[item]; !ok {
				ctx.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateResponseWithError(nil, -1, false, errors.New("dont have required roles")))
				return
			}
		}
		ctx.Next()
	}
}