package middleware

import (
	"net/http"

	"github.com/Arshia-Izadyar/Jabama-clone/src/api/helper"
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)

func Limiter() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(1, nil)
	return func(ctx *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, ctx.Writer, ctx.Request)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, helper.GenerateResponseWithError(nil, -1, false, err))
			return
		}
		ctx.Next()
	}
}
