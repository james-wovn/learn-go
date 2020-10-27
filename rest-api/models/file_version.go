package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type FileVersion struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time      `gorm:"type:timestamp without time zone;not null" json:"createdAt"`
	UpdatedAt  time.Time      `gorm:"type:timestamp without time zone;not null" json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"type:timestamp without time zone;index" json:"deletedAt"`
	File       File
	FileID     uint `gorm:"not null;index" json:"fileId"`
	User       User
	UserID     uint           `gorm:"not null;index" json:"userId"`
	Version    uint           `gorm:"type:serial" json:"version"`
	IsCurrent  bool           `gorm:"type:boolean;not null;default:true" json:"isCurrent"`
	Properties datatypes.JSON `gorm:"type:jsonb;not null;index:,type:gin" json:"properties"`
}
