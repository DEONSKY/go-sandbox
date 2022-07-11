package model

import (
	"time"

	"gorm.io/gorm"
)

//Book struct represents books table in database
type Project struct {
	ID              uint64         `gorm:"primary_key:auto_increment" json:"id"`
	Title           string         `gorm:"type:varchar(255)" json:"title"`
	Description     string         `gorm:"type:text" json:"description"`
	ProjectLeaderID uint64         `gorm:"not null" json:"-"`
	ProjectLeader   User           `gorm:"foreignkey:ProjectLeaderID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	Subjects        []Subject      `json:"-"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `json:"-"`
}
