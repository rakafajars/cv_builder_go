package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserID   uint   `gorm:"uniqueIndex;not null" json:"user_id"`
	FullName string `grom:"type:varchar(100)" json:"full_name"`
	Phone    string `gorm:"type:varchar(20)" json:"phone"`
	Address  string `gorm:"type:text" json:"address"`
	PhotoUrl string `gorm:"type:varchar(255)" json:"photo_url"`
	Summary  string `gorm:"type:text" json:"summary"`
}
