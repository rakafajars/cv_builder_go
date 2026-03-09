package models

type ProjectRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	TechStack   string `json:"tech_stack"`
}

func (r *ProjectRequest) ToModel(userID uint) Projects {
	return Projects{
		UserID:      userID,
		Title:       r.Title,
		Link:        r.Link,
		TechStack:   r.TechStack,
		Description: r.Description,
	}
}

type ProjectResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	TechStack   string `json:"tech_stack"`
}

func (r *ProjectRequest) ToResponse(id uint) ProjectResponse {
	return ProjectResponse{
		ID:          id,
		Title:       r.Title,
		Link:        r.Link,
		TechStack:   r.TechStack,
		Description: r.Description,
	}
}
