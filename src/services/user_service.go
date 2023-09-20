package services

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/Arshia-Izadyar/Jabama-clone/src/api/dto"
	"github.com/Arshia-Izadyar/Jabama-clone/src/common"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/constants"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/db"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/models"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/logger"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/service_errors"
)

type UserService struct {
	DB     *gorm.DB
	Logger logger.Logger
	Otp    *OtpService
	Token  *TokenService
}

func NewUserService(cfg *config.Config) *UserService {
	db := db.GetDB()
	logger := logger.NewLogger(cfg)
	otp := NewOtpService(cfg)
	tk := NewTokenService(cfg)
	return &UserService{
		DB:     db,
		Logger: logger,
		Otp:    otp,
		Token:  tk,
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

func (us *UserService) SendOtp(req *dto.OtpRequest) error {
	otp := common.GenerateOtp()
	err := us.Otp.SetOtp(req.PhoneNumber, otp)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) RegisterLoginByPhoneNumber(req *dto.RegisterLoginByPhone) (*dto.TokenDetail, error) {
	exists, err := us.checkByPhone(req.PhoneNumber)
	if err != nil {
		return nil, err
	}
	err = us.Otp.ValidateOtp(req.PhoneNumber, req.Otp)
	if err != nil {
		return nil, err
	}

	if !exists {
		user := &models.User{PhoneNumber: req.PhoneNumber, Username: req.PhoneNumber}
		bs, err := bcrypt.GenerateFromPassword([]byte("changeme"), bcrypt.MinCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(bs)
		roleId, err := us.getDefaultRole()
		if err != nil {
			return nil, err
		}
		tx := us.DB.Begin()
		err = tx.Model(&models.User{}).Create(user).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		userRole := &models.UserRole{RoleId: roleId, UserId: user.Id}
		err = tx.Create(userRole).Error
		if err != nil {
			return nil, err
		}
		tx.Commit()

		createdUser := &models.User{}

		err = us.DB.Model(&models.User{}).Where("username = ?", user.Username).Preload("UserRoles.Role").First(&createdUser).Error
		if err != nil {
			return nil, err
		}
		tDTO := &dto.TokenDto{
			UserId:   createdUser.Id,
			Username: createdUser.Username,
			Phone:    createdUser.PhoneNumber,
		}
		if len(createdUser.UserRoles) > 0 {
			for _, r := range createdUser.UserRoles {
				tDTO.Roles = append(tDTO.Roles, r.Role.Name)
			}
		}
		token, err := us.Token.GenerateToken(tDTO)
		if err != nil {
			return nil, err
		}
		return token, nil
	}
	var user models.User
	err = us.DB.Model(&models.User{}).Where("phone_number = ?", req.PhoneNumber).Preload("UserRoles.Role").First(&user).Error
	if err != nil {
		return nil, err
	}
	tDto := &dto.TokenDto{
		UserId:   user.Id,
		Username: user.Username,
		Phone:    user.PhoneNumber,
	}
	if len(user.UserRoles) > 0 {
		for _, r := range user.UserRoles {
			tDto.Roles = append(tDto.Roles, r.Role.Name)
		}
	}
	tk, err := us.Token.GenerateToken(tDto)
	if err != nil {
		return nil, err
	}
	return tk, nil

}

func (us *UserService) LoginByUserName(req *dto.LoginByUserName) (*dto.TokenDetail, error) {
	var user models.User
	err := us.DB.Model(&models.User{}).Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	tDto := &dto.TokenDto{
		UserId:   user.Id,
		Username: user.Username,
		Phone:    user.PhoneNumber,
	}
	if len(user.UserRoles) > 0 {
		for _, r := range user.UserRoles {
			tDto.Roles = append(tDto.Roles, r.Role.Name)
		}
	}

	tk, err := us.Token.GenerateToken(tDto)
	if err != nil {
		return nil, err
	}
	return tk, nil

}
