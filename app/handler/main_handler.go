package handler

import (
	"blog-api/interfaces"
	"blog-api/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MainHandler struct{}

func NewMainHandler() interfaces.MainHandlerInterface {
	return &MainHandler{}
}

func (h *MainHandler) HealthCheck() echo.HandlerFunc {
	return func(c echo.Context) error {
		response := response.NewApiResponse(c.Path())
		return c.JSON(http.StatusOK, response)
	}
}
