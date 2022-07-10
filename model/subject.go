package model

import (
	"time"

	"gorm.io/gorm"
)

//Book struct represents books table in database
type Subject struct {
	ID           uint64         `gorm:"primary_key:auto_increment" json:"id"`
	Title        string         `gorm:"type:varchar(255)" json:"title"`
	Description  string         `gorm:"type:text" json:"description"`
	RepoID       string         `gorm:"type:text" json:"repoId"`
	ProjectID    uint64         `gorm:"not null" json:"-"`
	Project      Project        `gorm:"foreignkey:ProjectID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"project"`
	Issues       []Issue        `json:"issues"`
	Stages       []Stage        `gorm:"foreignkey:id;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"stages"`
	User         []User         `gorm:"many2many:SubjectUser;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	TeamLeaderID uint64         `gorm:"not null" json:"-"`
	TeamLeader   User           `gorm:"foreignkey:TeamLeaderID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"teamLeader"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"-"`
}
