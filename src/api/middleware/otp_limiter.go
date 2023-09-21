package middleware

import (
	"errors"
	"net/http"
	"time"

	"github.com/Arshia-Izadyar/Jabama-clone/src/api/helper"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/limiter"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)


func OtpLimiter(cfg *config.Config) gin.HandlerFunc {
	var limiter = limiter.NewIpRateLimiter(rate.Every(cfg.Otp.Limiter*time.Second), 1)
	return func(ctx *gin.Context) {
		limiter := limiter.GetLimiter(ctx.ClientIP())
		if !limiter.Allow() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, helper.GenerateResponseWithError(nil, -1, false, errors.New("too many otp requests")))
			return
		}
		ctx.Next()
	}
}