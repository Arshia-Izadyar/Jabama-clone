package db

import (
	"fmt"
	"time"

	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var log = logger.NewLogger(config.GetConfig())
var SqlDb *gorm.DB

func InitDB(cfg *config.Config) error {
	var err error
	cnn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DbName,
	)
	SqlDb, err = gorm.Open(postgres.Open(cnn), &gorm.Config{})
	if err != nil {
		return err
	}

	db, err := SqlDb.DB()
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	db.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	db.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	db.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Minute)
	log.Info(logger.Postgres, logger.Startup, "postgres started", nil)
	return nil
}
func GetDB() *gorm.DB {
	return SqlDb
}

func CloseDB() {
	dataBase, err := SqlDb.DB()
	if err != nil {
		log.Fatal(logger.Postgres, logger.Close, err, nil)
	}
	err = dataBase.Close()
	if err != nil {
		log.Fatal(logger.Postgres, logger.Close, err, nil)
	}
}
