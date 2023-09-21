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
	Name                string           `json:"name"`
	Description         string           `json:"description"`
	Available           bool             `json:"available"`
	MinReserveTime      int              `json:"minReserveTime"`
	Cap                 int              `json:"cap"`
	AddedPricePerPerson float64          `json:"addedPricePerPerson"`
	PricePerNight       float64          `json:"pricePerNight"`
	City                CityResponse     `json:"city"`
	Province            ProvinceResponse `json:"province"`
}
