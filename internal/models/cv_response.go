package models

type CVResponse struct {
	Profile         *Profile         `json:"profile"`
	WorkExperiences []WorkExperience `json:"work_experiences"` // Pastikan nama struct-nya sesuai dengan yang ada di modelmu
	Educations      []Education      `json:"educations"`
	Skills          []Skills         `json:"skills"`
	Projects        []Projects       `json:"projects"`
}
