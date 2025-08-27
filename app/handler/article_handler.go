package handler

import (
	"blog-api/dto"
	"blog-api/interfaces"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	articleService interfaces.ArticleServiceInterface
}

func NewArticleHandler(
	articleService interfaces.ArticleServiceInterface,
) interfaces.ArticleHandlerInterface {
	return &ArticleHandler{
		articleService: articleService,
	}
}

func (h *ArticleHandler) GetList() echo.HandlerFunc {
	return func(c echo.Context) error {
		response := dto.NewApiResponse(c.Path())

		data, errs := h.articleService.GetList(c.Request().Context())
		if errs != nil {
			response.Errors = errs
			return c.JSON(http.StatusBadRequest, response)
		}

		response.Data = data

		return c.JSON(http.StatusOK, response)
	}
}

func (h *ArticleHandler) GetDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		response := dto.NewApiResponse(c.Path())

		var request dto.GetArticleDetailRequest
		var errs []error

		if err := c.Bind(&request); err != nil {
			errs = append(errs, err)
			response.Errors = errs
			return c.JSON(http.StatusBadRequest, response)
		}

		response.Request = request

		data, errs := h.articleService.GetDetail(c.Request().Context(), request)
		if errs != nil {
			response.Errors = errs
			return c.JSON(http.StatusBadRequest, response)
		}

		response.Data = data

		return c.JSON(http.StatusOK, response)
	}
}

func (h *ArticleHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		response := dto.NewApiResponse(c.Path())

		var request dto.CreateArticleRequest
		var errs []error
		if err := c.Bind(&request); err != nil {
			errs = append(errs, err)
			response.Errors = errs
			return c.JSON(http.StatusBadRequest, response)
		}
		response.Request = request

		data, errs := h.articleService.Create(c.Request().Context(), request)
		if errs != nil {
			response.Errors = errs
			return c.JSON(http.StatusBadRequest, response)
		}

		response.Data = data

		return c.JSON(http.StatusOK, response)
	}
}
