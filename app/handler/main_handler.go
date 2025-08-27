package handler

import (
	"blog-api/dto"
	"blog-api/interfaces"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type MainHandler struct{}

func NewMainHandler() interfaces.MainHandlerInterface {
	return &MainHandler{}
}

// @Summary      Check health
// @Description  Check health
// @Tags         main
// @Accept       json
// @Produce      json
// @Success      200  {object} dto.ApiResponse
// @Router       /api/healthy [get]
func (h *MainHandler) HealthCheck() echo.HandlerFunc {
	return func(c echo.Context) error {
		response := dto.NewApiResponse(c.Path())
		return c.JSON(http.StatusOK, response)
	}
}
