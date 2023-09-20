package services

import (
	"database/sql"

	"github.com/Arshia-Izadyar/Jabama-clone/src/api/dto"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/constants"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/db"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/models"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/logger"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/service_errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	DB     *gorm.DB
	Logger logger.Logger
}

func NewUserService(cfg *config.Config) *UserService {
	db := db.GetDB()
	logger := logger.NewLogger(cfg)
	return &UserService{
		DB:     db,
		Logger: logger,
	}
}

func (us *UserService) checkByEmail(email string) (bool, error) {
	var exists bool
	err := us.DB.Model(&models.User{}).Select("count(*) > 0").Where("email = ?", email).Find(&exists).Error
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (us *UserService) checkByUsername(username string) (bool, error) {
	var exists bool
	err := us.DB.Model(&models.User{}).Select("count(*) > 0").Where("username = ?", username).Find(&exists).Error
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (us *UserService) checkByPhone(number string) (bool, error) {
	var exists bool
	err := us.DB.Model(&models.User{}).Select("count(*) > 0").Where("phone_number = ?", number).Find(&exists).Error
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (us *UserService) getDefaultRole() (roleId int, err error) {
	if err := us.DB.Model(&models.Role{}).
		Select("id").
		Where("name = ?", constants.DefaultRoleName).
		First(&roleId).
		Error; err != nil {
		return -1, err
	}
	return roleId, nil
}

func (us *UserService) RegisterByUsername(req *dto.RegisterByUsername) error {
	user := &models.User{
		Username:  req.Username,
		Email:     sql.NullString{Valid: true, String: req.Email},
		FirstName: sql.NullString{Valid: true, String: req.FirstName},
		LastName:  sql.NullString{Valid: true, String: req.LastName},
	}
	exists, err := us.checkByUsername(req.Username)
	if err == nil && exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.UsernameExists}
	} else if err != nil {
		return err
	}

	exists, err = us.checkByEmail(req.Email)
	if err == nil && exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.EmailExists}
	} else if err != nil {
		return err
	}

	bs, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	user.Password = string(bs)

	roleId, err := us.getDefaultRole()
	if err != nil {
		return err
	}
	tx := us.DB.Begin()

	err = tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		us.Logger.Error(logger.Postgres, logger.Insert, err, nil)
		return err
	}
	userRole := &models.UserRole{UserId: user.Id, RoleId: roleId}
	err = tx.Create(&userRole).Error
	if err != nil {
		tx.Rollback()
		us.Logger.Error(logger.Postgres, logger.Insert, err, nil)
		return err
	}
	tx.Commit()
	return nil
}
