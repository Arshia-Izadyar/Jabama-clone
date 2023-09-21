package services

import (
	"context"
	"errors"

	"github.com/Arshia-Izadyar/Jabama-clone/src/api/dto"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/constants"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/db"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/models"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/logger"
	"gorm.io/gorm"
)

type UserWishListService struct {
	base BaseService[models.UserWishList, dto.CreateUserWishListRequest, dto.CreateUserWishListRequest, dto.UserWishListResponse]
}

func NewUserWishListService(cfg *config.Config) *UserWishListService {
	return &UserWishListService{
		base: BaseService[models.UserWishList, dto.CreateUserWishListRequest, dto.CreateUserWishListRequest, dto.UserWishListResponse]{
			DB:       db.GetDB(),
			Log:      logger.NewLogger(cfg),
			Preloads: []preload{},
		},
	}
}

func (uw *UserWishListService) GetById(ctx context.Context, id int) (*dto.UserWishListResponse, error) {
	model := &models.UserWishList{}
	userId := int64(ctx.Value(constants.UserIdKey).(float64))
	tx := uw.base.DB.WithContext(ctx).Begin()
	err := tx.Model(&model).Where("id = ?", id).Where("user_id = ?", userId).First(&model).Error
	if err != nil {
		tx.Rollback()

		uw.base.Log.Error(logger.Postgres, logger.Insert, err, nil)
		return nil, err
	}
	res := &dto.UserWishListResponse{
		// Id:          model.Id,
		UserId:      int(userId),
		ResidenceId: model.ResidenceId,
	}

	return res, nil
}

func (uw *UserWishListService) Create(ctx context.Context, req *dto.CreateUserWishListRequest) (*dto.UserWishListResponse, error) {
	req.UserId = int(ctx.Value(constants.UserIdKey).(float64))
	return uw.base.Create(ctx, req)
}

func (uw *UserWishListService) Delete(ctx context.Context, id int) error {
	model := models.UserWishList{}
	userId := int64(ctx.Value(constants.UserIdKey).(float64))

	tx := uw.base.DB.WithContext(ctx).Begin()
	err := tx.Model(&model).Where("id = ?", id).Where("user_id = ?", userId).First(&model).Error

	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		uw.base.Log.Error(logger.Postgres, logger.Get, errors.New("can't delete property category"), nil)
		return err
	}

	err = tx.Delete(&model).Error
	if err != nil {
		tx.Rollback()
		uw.base.Log.Error(logger.Postgres, logger.Delete, err, nil)
		return err
	}
	tx.Commit()
	return nil
}
