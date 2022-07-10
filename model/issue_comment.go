package model

import (
	"time"

	"gorm.io/gorm"
)

type IssueComment struct {
	ID        uint64         `gorm:"primary_key:auto_increment" json:"id"`
	Context   string         `gorm:"type:text" json:"context"`
	CreatorID uint64         `gorm:"not null" json:"-"`
	IssueID   uint64         `gorm:"not null" json:"-"`
	Creator   User           `gorm:"foreignkey:CreatorID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"creator"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
