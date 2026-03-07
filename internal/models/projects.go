package models

import "gorm.io/gorm"

type Projects struct {
	gorm.Model
	UserID      uint   `gorm:"not null" json:"user_id"`
	Title       string `gorm:"type:varchar(150)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Link        string `gorm:"type:varchar(255)" json:"link"`
	TechStack   string `gorm:"type:varchar(255)" json:"tech_stack"` // Bisa diisi string seperti "Flutter, Golang, Firebase"
}
