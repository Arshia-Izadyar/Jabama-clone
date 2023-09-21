package router

import (
	"github.com/Arshia-Izadyar/Jabama-clone/src/api/handler"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/gin-gonic/gin"
)

func CityRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCityHandler(cfg)
	r.POST("/", h.CreateCity)
	r.GET("/:id", h.GetById)
	r.PUT("/:id", h.UpdateCity)
	r.DELETE("/:id", h.DeleteCity)
	r.POST("/filter", h.GetByFilter)
}

func ProvinceRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewProvinceHandler(cfg)
	r.POST("/", h.CreateProvince)
	r.GET("/:id", h.GetById)
	r.PUT("/:id", h.UpdateProvince)
	r.DELETE("/:id", h.DeleteProvince)
	r.POST("/filter", h.GetByFilter)

}

func PropertyCategoryRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewPropertyCategoryHandler(cfg)
	r.POST("/", h.CreatePropertyCategory)
	r.GET("/:id", h.GetById)
	r.PUT("/:id", h.UpdatePropertyCategory)
	r.DELETE("/:id", h.DeletePropertyCategory)
	r.POST("/filter", h.GetByFilter)
}

func PropertyRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewPropertyHandler(cfg)
	r.POST("/", h.CreateProperty)
	r.GET("/:id", h.GetById)
	r.PUT("/:id", h.UpdateProperty)
	r.DELETE("/:id", h.DeleteProperty)
	r.POST("/filter", h.GetByFilter)

}
