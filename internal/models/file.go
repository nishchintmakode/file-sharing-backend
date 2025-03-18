package models

import (
	"time"

	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	UserID     uint      `gorm:"not null"`
	Filename   string    `gorm:"not null"`
	Size       int64     `gorm:"not null"`
	S3URL      string    `gorm:"not null"`
	UploadDate time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
