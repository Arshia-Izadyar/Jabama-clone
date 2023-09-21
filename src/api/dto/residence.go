package dto

/*
BaseModel
	Name                string  `gorm:"type:string;size:100;not null"`
	Description         string  `gorm:"type:string;size:1500"`
	Available           bool    `gorm:"default:true"`
	MinReserveTime      int     `gorm:"not null"`
	Cap                 int     `gorm:"not null"`
	AddedPricePerPerson float64 `gorm:"type:decimal(10,2);not null"`
	PricePerNight       float64 `gorm:"type:decimal(10,2);not null"`
	City                City    `gorm:"foreignKey:CityId;constraint:OnDelete:NO ACTION"`
	CityId              int
	Province            Province `gorm:"foreignKey:ProvinceId;constraint:OnDelete:NO ACTION"`
	ProvinceId          intson float64 `gorm:"type:decimal(10,2);not null"`
}
*/

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
	RoomCount    int    `json:"roomCount"`
	Type         string `json:"type"`
	BedType      string `json:"bedType"`
	HasShower    bool   `json:"hasShower"`
	HasBalcony   bool   `json:"hasBalcony"`
	HasPool      bool   `json:"hasPool"`
	HasGameRoom  bool   `json:"hasGameRoom"`
	HasFurniture bool   `json:"hasFurniture"`
}

type UpdateRoomTypeRequest struct {
	Type         string `json:"type"`
	BedType      string `json:"bedType"`
	HasShower    bool   `json:"hasShower"`
	HasBalcony   bool   `json:"hasBalcony"`
	HasPool      bool   `json:"hasPool"`
	HasGameRoom  bool   `json:"hasGameRoom"`
	HasFurniture bool   `json:"hasFurniture"`
}

type RoomTypeResponse struct {
	Id           int    `json:"id"`
	RoomCount    int    `json:"roomCount"`
	Type         string `json:"type"`
	BedType      string `json:"bedType"`
	HasShower    bool   `json:"hasShower"`
	HasBalcony   bool   `json:"hasBalcony"`
	HasPool      bool   `json:"hasPool"`
	HasGameRoom  bool   `json:"hasGameRoom"`
	HasFurniture bool   `json:"hasFurniture"`
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
