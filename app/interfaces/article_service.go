package interfaces

import (
	"blog-api/dto"

	"github.com/labstack/echo/v4"
)

type ArticleServiceInterface interface {
	GetList(ctx echo.Context) ([]dto.ArticleResponse, []error)
	GetDetail(ctx echo.Context, params dto.GetArticleDetailRequest) (*dto.ArticleResponse, []error)
	Create(ctx echo.Context, data dto.CreateArticleRequest) (*dto.ArticleResponse, []error)
}
