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
	Progress        uint32         `json:"progress"`
	SubjectID       uint64         `gorm:"not null" json:"-"`
	Subject         Subject        `gorm:"foreignkey:SubjectID;" json:"subject"`
	ParentIssueID   *uint64        `json:"p"`
	Status          uint8          `json:"status"`
	ChildIssues     []Issue        `gorm:"foreignkey:ParentIssueID;" json:"issues"`
	DependentIssues []Issue        `gorm:"many2many:DependentIssues;" json:"dependentIssues"`
	Comments        []IssueComment `gorm:"foreignkey:IssueID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"comments"`
	CreatorID       uint64         `gorm:"not null" json:"-"`
	Creator         User           `gorm:"foreignkey:CreatorID;" json:"creator"`
	AssignieID      *uint64        `json:"-"`
	Assignie        User           `gorm:"foreignkey:AssignieID;" json:"assignie"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `json:"-"`
}
