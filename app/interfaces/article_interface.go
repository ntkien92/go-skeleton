package interfaces

import (
	"blog-api/dto"
	"blog-api/model"
	"context"

	echo "github.com/labstack/echo/v4"
)

type ArticleRepositoryInterface interface {
	GetList(ctx context.Context, preloads []string) ([]model.Article, []error)
	GetDetail(ctx context.Context, params model.GetDetailArticleQueryParams) (*model.Article, []error)
	Create(ctx context.Context, data model.Article) (string, []error)
}

type ArticleServiceInterface interface {
	GetList(ctx context.Context) ([]dto.ArticleResponse, []error)
	GetDetail(ctx context.Context, params dto.GetArticleDetailRequest) (*dto.ArticleResponse, []error)
	Create(ctx context.Context, data dto.CreateArticleRequest) (*dto.ArticleResponse, []error)
}

type ArticleHandlerInterface interface {
	GetList() echo.HandlerFunc
	GetDetail() echo.HandlerFunc
	Create() echo.HandlerFunc
}
