package handler

import (
	"blog-api/interfaces"
	"blog-api/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ArticleHandler struct{}

func NewArticleHandler() interfaces.ArticleHandlerInterface {
	return &ArticleHandler{}
}

func (h *ArticleHandler) GetList() echo.HandlerFunc {
	return func(c echo.Context) error {
		response := response.NewApiResponse(c.Path())
		return c.JSON(http.StatusOK, response)
	}
}
