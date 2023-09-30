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

func RoomTypeRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewRoomTypeHandler(cfg)
	r.POST("/", h.CreateRoomType)
	r.GET("/:id", h.GetById)
	r.PUT("/:id", h.UpdateRoomType)
	r.DELETE("/:id", h.DeleteRoomType)
	r.POST("/filter", h.GetByFilter)
}

func ResidenceCommentRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewResidenceCommentHandler(cfg)
	r.POST("/", h.CreateResidenceComment)
	r.GET("/:id", h.GetById)
	r.PUT("/:id", h.UpdateResidenceComment)
	r.DELETE("/:id", h.DeleteResidenceComment)
	r.POST("/filter", h.GetByFilter)

}

func ResidencePropertyRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewResidencePropertyHandler(cfg)
	r.POST("/", h.CreateResidenceProperty)
	r.GET("/:id", h.GetById)
	r.PUT("/:id", h.UpdateResidenceProperty)
	r.DELETE("/:id", h.DeleteResidenceProperty)
	r.POST("/filter", h.GetByFilter)

}

func UserWishListRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewUserWishListHandler(cfg)
	r.POST("/", h.CreateUserWishList)
	r.GET("/:id", h.GetById)
	r.DELETE("/:id", h.DeleteUserWishList)
}

func ResidenceRateRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewResidenceRateHandler(cfg)
	r.POST("/", h.CreateResidenceRate)
	r.GET("/:id", h.GetById)
	r.PUT("/:id", h.UpdateResidenceRate)
	r.DELETE("/:id", h.DeleteResidenceRate)

}
