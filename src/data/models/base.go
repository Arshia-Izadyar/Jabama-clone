package models

type City struct {
	BaseModel
	Name      string `gorm:"type:string;size:30;not null;unique"`
	Provinces []Province
}

type Province struct {
	BaseModel
	Name   string `gorm:"type:string;size:30;not null;unique"`
	CityId int
	City   City `gorm:"foreignKey:CityId;constraint:onDelete:Cascade"`
}
