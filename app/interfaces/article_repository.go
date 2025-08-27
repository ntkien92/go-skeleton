package interfaces

import (
	"blog-api/model"
	"context"
)

type ArticleRepositoryInterface interface {
	GetList(ctx context.Context) ([]model.Article, []error)
	GetDetail(ctx context.Context, params model.GetDetailArticleQueryParams) (*model.Article, []error)
	Create(ctx context.Context, data model.Article) (string, []error)
}
