package dto

import (
	"blog-api/model"
	"time"
)

type ArticleResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewListArticleResponse(data []model.Article) []ArticleResponse {
	var result []ArticleResponse
	for _, article := range data {
		result = append(result, NewArticleDetailResponse(article))
	}

	return result
}

func NewArticleDetailResponse(data model.Article) ArticleResponse {
	return ArticleResponse{
		ID:        data.ID.String(),
		Title:     data.Title,
		Content:   data.Content,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

type CreateArticleRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type GetArticleDetailRequest struct {
	Id string `param:"id"`
}
