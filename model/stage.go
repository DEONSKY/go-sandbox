package model

import (
	"time"

	"gorm.io/gorm"
)

//Book struct represents books table in database
type Stage struct {
	ID          uint64         `gorm:"primary_key:auto_increment" json:"id"`
	Title       string         `gorm:"type:varchar(255)" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	ProjectId   uint64         `json:"-"`
	StartTime   time.Time      `json:"startTime"`
	EndTime     time.Time      `json:"endTime"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}
