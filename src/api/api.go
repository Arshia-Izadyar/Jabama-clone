package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/Arshia-Izadyar/Jabama-clone/src/api/middleware"
	"github.com/Arshia-Izadyar/Jabama-clone/src/api/router"
	"github.com/Arshia-Izadyar/Jabama-clone/src/api/validators"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/logger"
)

func Init(cfg *config.Config) {
	log := logger.NewLogger(cfg)
	r := gin.New()
	r.Use(gin.Recovery(), middleware.CustomLogger(log))
	r.Use(middleware.Limiter())
	r.Use(middleware.Cors(cfg))
	registerRoutes(r, cfg)
	registerValidators()

	log.Info(logger.General, logger.Startup, fmt.Sprintf("started listening on port %d", cfg.Server.Port), nil)
	err := r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		log.Fatal(logger.General, logger.Startup, err, nil)
		return
	}
}

func registerRoutes(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	// users
	users := v1.Group("/users")
	router.UserRouter(users, cfg)

	cities := v1.Group("/city", middleware.Authentication(cfg))
	router.CityRouter(cities, cfg)

	provinces := v1.Group("/province", middleware.Authentication(cfg))
	router.ProvinceRouter(provinces, cfg)

	residence := v1.Group("/residence", middleware.Authentication(cfg))
	router.ResidenceRouter(residence, cfg)

	// RoomTypeRouter

	room := v1.Group("/room", middleware.Authentication(cfg))
	router.RoomTypeRouter(room, cfg)

	comment := v1.Group("/comment", middleware.Authentication(cfg))
	router.ResidenceCommentRouter(comment, cfg)

	propertyCategory := v1.Group("/property-category", middleware.Authentication(cfg))
	router.PropertyCategoryRouter(propertyCategory, cfg)

	property := v1.Group("/property", middleware.Authentication(cfg))
	router.PropertyRouter(property, cfg)

	residenceProperty := v1.Group("/residence-property", middleware.Authentication(cfg))
	router.ResidencePropertyRouter(residenceProperty, cfg)

	wishList := v1.Group("/wishlist", middleware.Authentication(cfg))
	router.UserWishListRouter(wishList, cfg)

	rateRouter := v1.Group("/rate", middleware.Authentication(cfg))
	router.ResidenceRateRouter(rateRouter, cfg)

}

func registerValidators() {
	vld, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		vld.RegisterValidation("phone", validators.IranPhoneNumberValidator, true)
		vld.RegisterValidation("rate", validators.RateValidator, true)
	}
}
