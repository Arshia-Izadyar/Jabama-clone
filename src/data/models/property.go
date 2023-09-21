package models

type PropertyCategory struct {
	BaseModel
	Name       string     `gorm:"type:string;not null;unique"`
	Icon       string     `gorm:"type:string;not null;"`
	Properties []Property `gorm:"foreignKey:CategoryId"`
}

type Property struct {
	BaseModel
	Description string           `gorm:"type:string;size:1000;null"`
	Name        string           `gorm:"type:string;size:15;not null;unique"`
	Icon        string           `gorm:"type:string;size:1000;null"`
	Category    PropertyCategory `gorm:"foreignKey:CategoryId;constraint:onDelete:NO ACTION"`
	CategoryId  int
}
