package router

import (
	"github.com/Arshia-Izadyar/Jabama-clone/src/api/handler"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/gin-gonic/gin"
)

func ResidenceRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewResidenceHandler(cfg)
	r.POST("/", h.CreateResidence)
	r.GET("/:id", h.GetById)
	r.PUT("/:id", h.UpdateResidence)
	r.DELETE("/:id", h.DeleteResidence)
	r.POST("/filter", h.GetByFilter)

}
