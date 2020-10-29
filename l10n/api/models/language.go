package models

import (
	"gorm.io/gorm"
	"time"
)

type Language struct {
	Code      string         `gorm:"primaryKey;type:character(2)" json:"code"`
	CreatedAt time.Time      `gorm:"type:timestamp without time zone;not null" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"type:timestamp without time zone;not null" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp without time zone;index" json:"deletedAt"`
	Name      string         `gorm:"type:character varying;not null" json:"name"`
	UnitType  string         `gorm:"type:character(4);not null" json:"unitType"`
	UnitPrice uint           `json:"unitPrice not null"`
}
