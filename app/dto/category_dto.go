package dto

import (
	"blog-api/model"
	"time"
)

type CategoryResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewCategoryDetailResponse(data model.Category) CategoryResponse {
	return CategoryResponse{
		ID:        data.ID.String(),
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
