package services

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/Arshia-Izadyar/Jabama-clone/src/api/dto"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/constants"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/cache"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/logger"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/service_errors"
)

type OtpService struct {
	Logger logger.Logger
	Cfg    *config.Config
	Redis  *redis.Client
}

func NewOtpService(cfg *config.Config) *OtpService {
	l := logger.NewLogger(cfg)
	r := cache.GetRedis()
	return &OtpService{
		Logger: l,
		Cfg:    cfg,
		Redis:  r,
	}
}

func (os *OtpService) SetOtp(mobile, otp string) *service_errors.ServiceError {
	key := fmt.Sprintf("%s:%s", constants.DefaultRedisKey, mobile)
	value := &dto.OtpDto{
		Value: otp,
		Used:  false,
	}
	res, err := cache.Get[dto.OtpDto](key, os.Redis)
	if err == nil && res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	} else if err == nil && !res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpExists}
	}
	err = cache.Set[dto.OtpDto](key, *value, os.Cfg.Otp.ExpireTime)
	if err != nil {
		return &service_errors.ServiceError{EndUserMessage: "cant set otp"}
	}
	return nil
}

func (os *OtpService) ValidateOtp(phoneNumber, otp string) error {
	userOtp, err := cache.Get[dto.OtpDto](fmt.Sprintf("%s:%s", constants.DefaultRedisKey, phoneNumber), os.Redis)
	if err != nil {
		return err
	}
	if userOtp.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	} else if !userOtp.Used && userOtp.Value != otp {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpInvalid}
	} else if !userOtp.Used && userOtp.Value == otp {
		userOtp.Used = true
		err = cache.Set[dto.OtpDto](fmt.Sprintf("%s:%s", constants.DefaultRedisKey, phoneNumber), *userOtp, os.Cfg.Otp.ExpireTime*time.Minute)
		if err != nil {
			return err
		}
	}
	return nil
}
