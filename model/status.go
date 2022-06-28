package model

//Book struct represents books table in database
type Status struct {
	ID          uint32 `gorm:"primary_key:auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(16)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	HexCode     string `gorm:"type:varchar(6)" json:"hexCode"`
}
