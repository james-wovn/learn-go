package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID                  uint           `gorm:"primaryKey" json:"id"`
	CreatedAt           time.Time      `gorm:"type:timestamp without time zone;not null" json:"createdAt"`
	UpdatedAt           time.Time      `gorm:"type:timestamp without time zone;not null" json:"updatedAt"`
	DeletedAt           gorm.DeletedAt `gorm:"type:timestamp without time zone;index" json:"deletedAt"`
	DisplayLanguage     Language       `gorm:"foreignKey:DisplayLanguageCode"`
	DisplayLanguageCode string         `gorm:"type:character(2);not null;default:en;index" json:"displayLanguage"`
	FileLanguage        Language       `gorm:"foreignKey:FileLanguageCode"`
	FileLanguageCode    string         `gorm:"type:character(2);not null;default:en;index" json:"fileLanguage"`
	Email               string         `gorm:"type:character varying;not null" json:"email"`
	Password            string         `gorm:"type:character varying;not null" json:"password"`
	Properties          datatypes.JSON `gorm:"type:jsonb;not null;index:,type:gin" json:"properties"`
}
