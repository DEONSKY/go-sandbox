package model

import (
	"time"

	"gorm.io/gorm"
)

//Book struct represents books table in database
type Issue struct {
	ID              uint64         `gorm:"primary_key:auto_increment" json:"id"`
	Title           string         `gorm:"type:varchar(255)" json:"title"`
	Description     string         `gorm:"type:text" json:"description"`
	IssueForeignId  string         `gorm:"type:text" json:"issueForeignId"`
	TargetTime      uint32         `json:"targetTime"`
	SpendingTime    uint32         `json:"spendingTime"`
	Progress        uint8          `json:"progress"`
	SubjectID       uint64         `gorm:"not null" json:"-"`
	Subject         Subject        `gorm:"foreignkey:SubjectID;" json:"-"`
	ParentIssueID   *uint64        `json:"parentIssueID"`
	StatusID        uint8          `gorm:"not null;default:1" json:"status"`
	ChildIssues     []Issue        `gorm:"foreignkey:ParentIssueID;" json:"-"`
	DependentIssues []Issue        `gorm:"many2many:DependentIssues;" json:"-"`
	Comments        []IssueComment `gorm:"foreignkey:ID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	ReporterID      uint64         `gorm:"not null" json:"-"`
	Reporter        User           `gorm:"foreignkey:ReporterID;" json:"-"`
	AssignieID      *uint64        `json:"-"`
	Assignie        User           `gorm:"foreignkey:AssignieID;" json:"-"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `json:"-"`
}
