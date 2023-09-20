package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/Arshia-Izadyar/Jabama-clone/src/api/router"
	"github.com/Arshia-Izadyar/Jabama-clone/src/api/validators"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/logger"
)

func Init(cfg *config.Config) {
	log := logger.NewLogger(cfg)
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())

	registerRoutes(r, cfg)

	log.Info(logger.General, logger.Startup, fmt.Sprintf("started listening on posrt %d", cfg.Server.Port), nil)
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
}

func registerValidators() {
	vld, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		vld.RegisterValidation("phone", validators.IranPhoneNumberValidator, true)
	}
}
