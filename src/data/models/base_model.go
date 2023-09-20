package models

import (
	"database/sql"
	"time"

	"github.com/Arshia-Izadyar/Jabama-clone/src/constants"
	"gorm.io/gorm"
)

type BaseModel struct {
	Id        int          `gorm:"primarykey"`
	CreatedAt time.Time    `gorm:"type:TIMESTAMP with time zone;not null"`
	UpdatedAt sql.NullTime `gorm:"type:TIMESTAMP with time zone;null"`

	CreatedBy int           `gorm:"not null"`
	UpdatedBy sql.NullInt64 `gorm:"null"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	value := tx.Statement.Context.Value(constants.UserIdKey)
	userId := -1
	if value != nil {
		userId = int(value.(float64))
	}

	b.CreatedBy = userId
	b.CreatedAt = time.Now()
	return nil
}

func (b *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	value := tx.Statement.Context.Value(constants.UserIdKey)
	userId := sql.NullInt64{Valid: true, Int64: -1}
	if value != nil {
		userId = sql.NullInt64{Valid: true, Int64: int64(value.(float64))}
	}
	b.UpdatedAt = sql.NullTime{Valid: true, Time: time.Now()}
	b.UpdatedBy = userId
	return nil
}
