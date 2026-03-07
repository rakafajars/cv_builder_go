package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email       string            `gorm:"uniqueIndex;not null" json:"email"`
	Password    string            `gorm:"not null" json:"-"`
	Profile     Profile           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"profile"`
	Experiences []WorkExperiences `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"experiences"`
	Education   []Education       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"educations"`
	Skills      []Skills          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"skills"`
	Projects    []Projects        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"projects"`
}
