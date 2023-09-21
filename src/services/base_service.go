package services

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Arshia-Izadyar/Jabama-clone/src/common"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/constants"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/db"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/models"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/logger"
	"gorm.io/gorm"
)

type preload struct {
	name string
}

type BaseService[T, Tu, Tc, Tr any] struct {
	DB *gorm.DB
	Log logger.Logger
	Preloads []preload
}

func NewBaseService[T, Tu, Tc, Tr any](cfg *config.Config) *BaseService[T, Tu, Tc, Tr] {
	db := db.GetDB()
	log := logger.NewLogger(cfg)
	return &BaseService[T, Tu, Tc, Tr]{
		DB:  db,
		Log: log,
		Preloads: []preload{},
	}
}

func LoadPreloads(db *gorm.DB, preload []preload) *gorm.DB{
	for _, i := range preload {
		err := db.Preload(i.name).Error
		if err != nil {
			panic(err)
		} else {
			db = db.Preload(i.name)
		}
	}
	return db


}


func (bs *BaseService[T, Tu, Tc, Tr]) GetById(ctx *context.Context, id int) (*Tr, error) {
	model := new(T)
	db := LoadPreloads(bs.DB, bs.Preloads)
	err := db.Model(&model).Where("id = ?", id).First(&model).Error
	if err != nil {
		return nil, err
	}
	res, err := common.TypeConvert[Tr](model)
	if err != nil {
		return nil,err
	}
	return res, nil
}

func (bs *BaseService[T, Tu, Tc, Tr]) Create(ctx context.Context, req *Tc) (*Tr, error) {
	model, err := common.TypeConvert[T](req)
	if err != nil {
		return nil, err
	}
	tx := bs.DB.WithContext(ctx).Begin()
	err = tx.Create(model).Error
	if err != nil {
		tx.Rollback()
		bs.Log.Error(logger.Postgres, logger.Insert, err, nil)
		return nil, err
	}
	tx.Commit()
	res, err := common.TypeConvert[models.BaseModel](model)
	if err != nil {
		return nil, err
	}
	return bs.GetById(&ctx,res.Id)
}

func (bs *BaseService[T, Tu, Tc, Tr]) Delete(ctx context.Context, id int) (error) {
	model := new(T)
	tx := bs.DB.WithContext(ctx).Begin()
	err := tx.First(&model, id).Error
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		bs.Log.Error(logger.Postgres, logger.Get, errors.New("can't delete property category"), nil)
		return err
	}

	err = tx.Delete(&model).Error
	if err != nil {
		tx.Rollback()
		bs.Log.Error(logger.Postgres, logger.Delete, err, nil)
		return err
	}
	tx.Commit()
	return nil
}


func (bs *BaseService[T, Tu, Tc, Tr]) Update(ctx context.Context, req *Tu, id int) (*Tr, error) {
	updateMap , err := common.TypeConvert[map[string]interface{}](req)
	if err != nil {
		return nil, err
	}
	snakeMap := map[string]interface{}{}
	for k, v := range *updateMap{
		snakeMap[common.ConvertToSnakeCase(k)] = v
	}
	snakeMap["updated_at"] = &sql.NullTime{Valid: true, Time: time.Now()}
	snakeMap["updated_by"] = &sql.NullInt64{Valid: true, Int64: int64(ctx.Value(constants.UserIdKey).(float64))}
	model := new(T)

	tx := bs.DB.WithContext(ctx).Begin()
	err = tx.Model(&model).Where("id = ?", id).Updates(snakeMap).Error
	if err != nil {
		tx.Rollback()
		bs.Log.Error(logger.Postgres, logger.Update, err, nil)
		return nil, err
	}
	tx.Commit()
	return bs.GetById(&ctx, id)
}