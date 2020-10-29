package models

import (
	"gorm.io/gorm"
	"time"
)

type ProjectAccessPermission struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"type:timestamp without time zone;not null" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"type:timestamp without time zone;not null" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp without time zone;index" json:"deletedAt"`
	Project   Project
	ProjectID uint `gorm:"not null;index" json:"projectId"`
	User      User
	UserID    uint `gorm:"not null;index" json:"userId"`
}
