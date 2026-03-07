package models

import (
	"time"

	"gorm.io/gorm"
)

type Education struct {
	gorm.Model
	UserID       uint       `gorm:"not null" json:"user_id"`
	Institution  string     `gorm:"type:varchar(150)" json:"institution"`
	Degree       string     `gorm:"type:varchar(100)" json:"degree"`
	FieldOfStudy string     `gorm:"type:varchar(100)" json:"field_of_study"`
	StartDate    time.Time  `gorm:"type:date" json:"start_date"`
	EndDate      *time.Time `gorm:"type:date" json:"end_date"`
	GPA          float64    `json:"gpa"`
}
