package models

import "database/sql"

type User struct {
	BaseModel
	Username     string         `gorm:"type:string;not null;unique;size:100"`
	Email        sql.NullString `gorm:"type:string;null;unique;size:150"`
	PhoneNumber  string         `gorm:"type:string;null;unique;size:15"`
	FirstName    sql.NullString `gorm:"type:string;null;size:30"`
	LastName     sql.NullString `gorm:"type:string;null;size:35"`
	Activated    bool           `gorm:"default:false"`
	Password     string         `gorm:"type:string;not null;size:64"`
	UserRoles    []UserRole
	UserWishList []UserWishList
}
type UserWishList struct {
	BaseModel
	User        User      `gorm:"foreignKey:UserId;constraint:onUpdate:NO ACTION;onDelete:NO ACTION"`
	Residence   Residence `gorm:"foreignKey:ResidenceId;constraint:OnDelete:CASCADE"`
	UserId      int
	ResidenceId int `gorm:"unique"`
}

type Role struct {
	BaseModel
	Name      string `gorm:"type:string;not null;unique"`
	UserRoles []UserRole
}

type UserRole struct {
	UserId int
	RoleId int
	User   User `gorm:"foreignKey:UserId"`
	Role   Role `gorm:"foreignKey:RoleId"`
}
