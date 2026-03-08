package models

import "time"

// WorkExperienceRequest adalah request body untuk Create dan Update Work Experience
type WorkExperienceRequest struct {
	CompanyName string     `json:"company_name"`
	Position    string     `json:"position"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	IsCurrent   bool       `json:"is_current"`
	Description string     `json:"description"`
}

// ToModel mengkonversi request menjadi model WorkExperience
func (r *WorkExperienceRequest) ToModel(userID uint) WorkExperience {
	return WorkExperience{
		UserID:      userID,
		CompanyName: r.CompanyName,
		Position:    r.Position,
		StartDate:   r.StartDate,
		EndDate:     r.EndDate,
		IsCurrent:   r.IsCurrent,
		Description: r.Description,
	}
}

// WorkExperienceResponse adalah response untuk Update Work Experience
type WorkExperienceResponse struct {
	ID          uint       `json:"id"`
	CompanyName string     `json:"company_name"`
	Position    string     `json:"position"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	IsCurrent   bool       `json:"is_current"`
	Description string     `json:"description"`
}

// ToResponse mengkonversi request menjadi response dengan ID
func (r *WorkExperienceRequest) ToResponse(id uint) WorkExperienceResponse {
	return WorkExperienceResponse{
		ID:          id,
		CompanyName: r.CompanyName,
		Position:    r.Position,
		StartDate:   r.StartDate,
		EndDate:     r.EndDate,
		IsCurrent:   r.IsCurrent,
		Description: r.Description,
	}
}
