package models

import "time"

type EducationRequest struct {
	Institution  string     `json:"institution"`
	Degree       string     `json:"degree"`
	FieldOfStudy string     `json:"field_of_study"`
	StartDate    time.Time  `json:"start_date"`
	EndDate      *time.Time `json:"end_date"`
	Gpa          float64    `json:"gpa"`
	Description  string     `json:"description"`
	IsCurrent    bool       `json:"is_current"`
}

func (r *EducationRequest) ToModel(userID uint) Education {
	return Education{
		UserID:       userID,
		Degree:       r.Degree,
		FieldOfStudy: r.FieldOfStudy,
		StartDate:    r.StartDate,
		EndDate:      r.EndDate,
		IsCurrent:    r.IsCurrent,
		GPA:          r.Gpa,
		Institution:  r.Institution,
	}
}

type EducationResponse struct {
	ID           uint       `json:"id"`
	Institution  string     `json:"institution"`
	Degree       string     `json:"degree"`
	FieldOfStudy string     `json:"field_of_study"`
	StartDate    time.Time  `json:"start_date"`
	EndDate      *time.Time `json:"end_date"`
	Gpa          float64    `json:"gpa"`
	Description  string     `json:"description"`
	IsCurrent    bool       `json:"is_current"`
}

func (r *EducationRequest) ToResponse(id uint) EducationResponse {
	return EducationResponse{
		ID:           id,
		Degree:       r.Degree,
		FieldOfStudy: r.FieldOfStudy,
		StartDate:    r.StartDate,
		EndDate:      r.EndDate,
		IsCurrent:    r.IsCurrent,
		Institution:  r.Institution,
		Gpa:          r.Gpa,
		Description:  r.Description,
	}
}
