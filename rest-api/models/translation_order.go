package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type TranslationOrder struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	CreatedAt       time.Time      `gorm:"type:timestamp without time zone;not null" json:"createdAt"`
	UpdatedAt       time.Time      `gorm:"type:timestamp without time zone;not null" json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"type:timestamp without time zone;index" json:"deletedAt"`
	FileVersion     FileVersion
	FileVersionID   uint `gorm:"not null;index" json:"fileVersionId"`
	User            User
	UserID          uint           `gorm:"not null;index" json:"userId"`
	Subtotal        uint           `gorm:"type:integer;not null" json:"subtotal"`
	Total           uint           `gorm:"type:integer;not null" json:"total"`
	TranslationJobs datatypes.JSON `gorm:"type:jsonb;not null;index:,type:gin" json:"translationJobs"`
}
