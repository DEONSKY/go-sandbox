package model

import "time"

//Book struct represents books table in database
type Project struct {
	ID          uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Title       string    `gorm:"type:varchar(255)" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Subjects    []Subject `json:"subjects"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
