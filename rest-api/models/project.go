package models

import (
	"gorm.io/gorm"
	"time"
)

// TODO: add `memsource_project_id`
type Project struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"type:timestamp without time zone;not null" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"type:timestamp without time zone;not null" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp without time zone;index" json:"deletedAt"`
	User      User
	UserID    uint   `gorm:"not null; index" json:"userId"`
	Name      string `gorm:"type:character varying;not null" json:"name"`
}
