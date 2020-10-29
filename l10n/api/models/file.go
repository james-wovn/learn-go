package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type File struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time      `gorm:"type:timestamp without time zone;not null" json:"createdAt"`
	UpdatedAt  time.Time      `gorm:"type:timestamp without time zone;not null" json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"type:timestamp without time zone;index" json:"deletedAt"`
	Folder     Folder
	FolderID   uint `gorm:"not null;index" json:"folderId"`
	Project    Project
	ProjectID  uint           `gorm:"not null;index" json:"projectId"`
	Properties datatypes.JSON `gorm:"type:jsonb;not null;index:,type:gin" json:"properties"`
}
