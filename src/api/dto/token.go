package dto

type TokenDetail struct {
	AccessToken            string `json:"accessToken"`
	RefreshToken           string `json:"refreshToken"`
	AccessTokenExpireTime  int64  `json:"accessTokenExpireTime"`
	RefreshTokenExpireTime int64  `json:"refreshTokenExpireTime"`
}

type TokenDto struct {
	UserId   int
	Username string
	Phone    string
	Roles    []string
}

type RefreshTokenDTO struct {
	RefreshToken string `json:"refreshToken"`
}
