package models

type SkillRequest struct {
	Name     string `json:"name"`
	Level    string `json:"level"`
	Category string `json:"category"`
}

func (r *SkillRequest) ToModel(userID uint) Skills {
	return Skills{
		UserID:   userID,
		Name:     r.Name,
		Level:    r.Level,
		Category: r.Category,
	}
}

type SkillResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Level    string `json:"level"`
	Category string `json:"category"`
}

func (r *SkillRequest) ToResponse(id uint) SkillResponse {
	return SkillResponse{
		ID:       id,
		Name:     r.Name,
		Level:    r.Level,
		Category: r.Category,
	}
}
