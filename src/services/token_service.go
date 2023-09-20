package services

import (
	"time"

	"github.com/Arshia-Izadyar/Jabama-clone/src/api/dto"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/constants"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/logger"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/service_errors"
	"github.com/golang-jwt/jwt/v4"
)

type TokenService struct {
	cfg    *config.Config
	logger logger.Logger
}

func NewTokenService(cfg *config.Config) *TokenService {
	l := logger.NewLogger(cfg)
	return &TokenService{
		cfg:    cfg,
		logger: l,
	}
}

func (ts *TokenService) GenerateToken(req *dto.TokenDto) (*dto.TokenDetail, error) {
	accessTokeDetail := dto.TokenDetail{}
	accessTokeDetail.AccessTokenExpireTime = time.Now().Add(ts.cfg.JWT.AccessTokenExpireDuration * time.Minute).Unix()
	accessTokeDetail.RefreshTokenExpireTime = time.Now().Add(ts.cfg.JWT.RefreshTokenExpireDuration * time.Minute).Unix()

	atClaims := jwt.MapClaims{}
	atClaims[constants.UserIdKey] = req.UserId
	atClaims[constants.PhoneKey] = req.Phone
	atClaims[constants.UsernameKey] = req.Username
	atClaims[constants.ExpKey] = accessTokeDetail.AccessTokenExpireTime
	atClaims[constants.AccessType] = true
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	var err error
	accessTokeDetail.AccessToken, err = tk.SignedString([]byte(ts.cfg.JWT.Secret))
	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims[constants.UserIdKey] = req.UserId
	rtClaims[constants.ExpKey] = accessTokeDetail.AccessTokenExpireTime

	rTk := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	accessTokeDetail.RefreshToken, err = rTk.SignedString([]byte(ts.cfg.JWT.Secret))
	if err != nil {
		return nil, err
	}
	return &accessTokeDetail, nil
}

func (ts *TokenService) validateToken(token string) (*jwt.Token, error) {
	tk, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, &service_errors.ServiceError{EndUserMessage: service_errors.TokenInvalid}
		}
		return []byte(ts.cfg.JWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	// TODO: blacklist check
	if claims, ok := tk.Claims.(jwt.MapClaims); ok && tk.Valid {
		expTime := time.Unix(int64(claims[constants.ExpKey].(float64)), 0)
		timeNow := time.Now()
		if timeNow.After(expTime) {
			return nil, &service_errors.ServiceError{EndUserMessage: service_errors.TokenExpired}
		}
	}
	return tk, nil
}
