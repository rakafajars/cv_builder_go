package models

import (
	"time"

	"gorm.io/gorm"
)

type WorkExperiences struct {
	gorm.Model
	UserID      uint       `gorm:"not null" json:"user_id"`
	CompanyName string     `gorm:"type:varchar(100)" json:"company_name"`
	Position    string     `gorm:"type:varchar(100)" json:"position"`
	StartDate   time.Time  `gorm:"type:date" json:"start_date"`
	EndDate     *time.Time `gorm:"type:date" json:"end_date"`
	IsCurrent   bool       `gorm:"default:false" json:"is_current"`
	Description string     `gorm:"type:text" json:"description"`
}
