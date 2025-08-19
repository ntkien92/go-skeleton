package interfaces

import (
	"blog-api/model"
	"context"
)

type ArticleRepositoryInterface interface {
	GetList(ctx context.Context) ([]model.Article, error)
}
