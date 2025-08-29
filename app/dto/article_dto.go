package dto

import (
	"blog-api/model"
	"time"
)

type ArticleResponse struct {
	ID        string            `json:"id"`
	Title     string            `json:"title"`
	Content   string            `json:"content"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt"`
	Category  *CategoryResponse `json:"category,omitempty"`
}

func NewListArticleResponse(data []model.Article) []ArticleResponse {
	var result []ArticleResponse
	for _, article := range data {
		result = append(result, NewArticleDetailResponse(article))
	}

	return result
}

func NewArticleDetailResponse(data model.Article) ArticleResponse {
	article := ArticleResponse{
		ID:        data.ID.String(),
		Title:     data.Title,
		Content:   data.Content,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	if data.Category != nil {
		category := NewCategoryDetailResponse(*data.Category)
		article.Category = &category
	}

	return article
}

type GetListArticleRequest struct {
	PagingRequest
}

type CreateArticleRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type GetArticleDetailRequest struct {
	Id string `param:"id"`
}
