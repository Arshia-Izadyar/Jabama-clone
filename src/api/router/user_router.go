package router

import (
	"github.com/Arshia-Izadyar/Jabama-clone/src/api/handler"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewUserHandler(cfg)
	r.POST("/username/signup", h.CreateUserByUsername)
	r.POST("/username/login", h.LoginByUsername)
	r.POST("/phone/signup", h.RegisterLoginByPhoneNumber)
	r.POST("/phone/login", h.RegisterLoginByPhoneNumber)
	r.POST("/otp", h.GetOtp)
	r.POST("/refresh", h.RefreshToken)
}
