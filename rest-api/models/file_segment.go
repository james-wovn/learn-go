package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type FileSegment struct {
	ID                      uint           `gorm:"primaryKey" json:"id"`
	CreatedAt               time.Time      `gorm:"type:timestamp without time zone;not null" json:"createdAt"`
	UpdatedAt               time.Time      `gorm:"type:timestamp without time zone;not null" json:"updatedAt"`
	DeletedAt               gorm.DeletedAt `gorm:"type:timestamp without time zone;index" json:"deletedAt"`
	FileVersion             FileVersion
	FileVersionID           uint           `gorm:"not null;index" json:"fileVersionId"`
	SourceLanguage          Language       `gorm:foreignKey;reference:SourceLanguageCode`
	SourceLanguageCode      string         `gorm:"type:character(2);not null;index" json:"sourceLanguageCode"`
	SourceText              string         `gorm:"not null" json:"sourceText"`
	SourceLanguageUnitCount uint           `gorm:"type:integer;not null" json:"sourceLanguageUnitCount"`
	Translation             datatypes.JSON `gorm:"type:jsonb;not null;index:,type:gin" json:"translation"`
}
