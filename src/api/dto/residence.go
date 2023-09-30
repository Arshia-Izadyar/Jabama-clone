package dto

type CreateResidenceRequest struct {
	Name                string  `json:"name"`
	Description         string  `json:"description"`
	Available           bool    `json:"available"`
	MinReserveTime      int     `json:"minReserveTime"`
	Cap                 int     `json:"cap"`
	AddedPricePerPerson float64 `json:"addedPricePerPerson"`
	PricePerNight       float64 `json:"pricePerNight"`
	CityId              int     `json:"cityId"`
	ProvinceId          int     `json:"provinceId"`
	RoomTypeId          int     `json:"roomTypeId"`
	OwnerId             int     `json:"ownerId"`
}

type UpdateResidenceRequest struct {
	Name                string  `json:"name"`
	Description         string  `json:"description"`
	Available           bool    `json:"available"`
	MinReserveTime      int     `json:"minReserveTime"`
	AddedPricePerPerson float64 `json:"addedPricePerPerson"`
	PricePerNight       float64 `json:"pricePerNight"`
}

type ResidenceResponse struct {
	Id                  int
	Name                string                      `json:"name"`
	Description         string                      `json:"description"`
	Available           bool                        `json:"available"`
	MinReserveTime      int                         `json:"minReserveTime"`
	Cap                 int                         `json:"cap"`
	AddedPricePerPerson float64                     `json:"addedPricePerPerson"`
	PricePerNight       float64                     `json:"pricePerNight"`
	OwnerId             int                         `json:"ownerId"`
	City                CityResponse                `json:"city"`
	Province            ProvinceResponse            `json:"province"`
	RoomType            RoomTypeResponse            `json:"roomType"`
	ResidenceComment    []ResidenceCommentResponse  `json:"residenceComment"`
	ResidenceProperties []ResidencePropertyResponse `json:"residenceProperties"`
}

type CreateRoomTypeRequest struct {
	RoomCount int    `json:"roomCount"`
	Type      string `json:"type"`
	BedType   string `json:"bedType"`
}

type UpdateRoomTypeRequest struct {
	Type      string `json:"type"`
	BedType   string `json:"bedType"`
	HasShower bool   `json:"hasShower"`
}

type RoomTypeResponse struct {
	Id        int    `json:"id" binding:"required"`
	RoomCount int    `json:"roomCount" binding:"required"`
	Type      string `json:"type" binding:"required"`
	BedType   string `json:"bedType" binding:"required"`
}

type CreateResidenceCommentRequest struct {
	UserId      int    `json:"userId"`
	ResidenceId int    `json:"residenceId"`
	Message     string `json:"message" binding:"max=1000"`
}

type UpdateResidenceCommentRequest struct {
	Message string `json:"message" binding:"max=1000"`
}

type ResidenceCommentResponse struct {
	Id      int         `json:"id"`
	Message string      `json:"message"`
	User    UsrResponse `json:"user"`
}
type UsrResponse struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phoneNumber"`
}

type CreateResidencePropertyRequest struct {
	PropertyId  int    `json:"propertyId"`
	ResidenceId int    `json:"residenceId"`
	Value       string `json:"value"`
}

type UpdateResidencePropertyRequest struct {
	Value string `json:"value"`
}

type ResidencePropertyResponse struct {
	Id          int              `json:"id"`
	Property    PropertyResponse `json:"property"`
	ResidenceId int              `json:"residenceId"`
	Value       string           `json:"value"`
}

type CreateUserWishListRequest struct {
	UserId      int `json:"userId"`
	ResidenceId int `json:"residenceId"`
}

type UserWishListResponse struct {
	// Id          int `json:"id"`
	UserId      int `json:"userId"`
	ResidenceId int `json:"residenceId"`
}

type CreateResidenceRateRequest struct {
	Rate        int `json:"rate"`
	ResidenceId int `json:"residenceId"`
	UserId      int `json:"userId"`
}

type UpdateResidenceRateRequest struct {
	Rate        int `json:"rate" binding:"required"`
	ResidenceId int `json:"residenceId" binding:"required"`
	UserId      int `json:"userId" binding:"required"`
}

type ResidenceRateResponse struct {
	Rate        int          `json:"rate"`
	ResidenceId int          `json:"residenceId"`
	UserId      UserResponse `json:"userId"`
}
