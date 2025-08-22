package interfaces

import (
	"blog-api/model"

	"github.com/labstack/echo/v4"
)

type ArticleRepositoryInterface interface {
	GetList(ctx echo.Context) ([]model.Article, []error)
	GetDetail(ctx echo.Context, params model.GetDetailArticleQueryParams) (*model.Article, []error)
	Create(ctx echo.Context, data model.Article) (string, []error)
}
