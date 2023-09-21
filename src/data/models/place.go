package models

/*
Residence {
	Type
	Location
	price
	Available
	Description
	Rate
	Cap
	TODO:Properties
	file
}
*/

type Residence struct {
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
	ProvinceId          int
	ResidenceRoom       []ResidenceRoom
	ResidenceComment    []ResidenceComment
}

type RoomType struct {
	BaseModel
	RoomCount    int    `gorm:"not null"`
	Type         string `gorm:"type:string;not null"`
	BedType      string `gorm:"type:string;not null"`
	HasShower    bool   `gorm:"default:false"`
	HasBalcony   bool   `gorm:"default:false"`
	HasPool      bool   `gorm:"default:false"`
	HasGameRoom  bool   `gorm:"default:false"`
	HasFurniture bool   `gorm:"default:false"`
}

type ResidenceRoom struct {
	BaseModel
	Residence   Residence `gorm:"foreignKey:ResidenceId;constraint:OnDelete:CASCADE"`
	ResidenceId int
	RoomType    RoomType `gorm:"foreignKey:RoomTypeId;constraint:OnDelete:NO ACTION"`
	RoomTypeId  int
}

type ResidenceComment struct {
	BaseModel
	Residence   Residence `gorm:"foreignKey:ResidenceId;constraint:OnDelete:CASCADE"`
	ResidenceId int
	Message     string `gorm:"size:1000;type:string;not null"`
	User        User   `gorm:"foreignKey:UserId;constraint:onUpdate:NO ACTION;onDelete:NO ACTION"`
	UserId      int
}
