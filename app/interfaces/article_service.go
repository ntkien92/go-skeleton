package interfaces

import (
	"blog-api/dto"
	"context"
)

type ArticleServiceInterface interface {
	GetList(ctx context.Context) ([]dto.ArticleResponse, []error)
	GetDetail(ctx context.Context, params dto.GetArticleDetailRequest) (*dto.ArticleResponse, []error)
	Create(ctx context.Context, data dto.CreateArticleRequest) (*dto.ArticleResponse, []error)
}
