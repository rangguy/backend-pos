package models

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UUID      uuid.UUID `gorm:"type:uuid;not null"`
	Code      string    `gorm:"type:varchar(100);unique;not null"`
	Name      string    `gorm:"type:varchar(255);not null"`
	PriceBuy  uint      `gorm:"type:uint;not null"`
	PriceSale uint      `gorm:"type:uint;not null"`
	Stock     uint      `gorm:"type:uint;not null"`
	Unit      string    `gorm:"type:varchar(100);not null"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
