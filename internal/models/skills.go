package models

import "gorm.io/gorm"

type Skills struct {
	gorm.Model
	UserID   uint   `gorm:"not null" json:"user_id"`
	Name     string `gorm:"type:varchar(50)" json:"name"`
	Level    string `gorm:"type:varchar(50)" json:"level"`
	Category string `gorm:"type:varchar(100)" json:"category"`
}
