package handler

import (
	"blog-api/dto"
	"blog-api/perrors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func errorResponse(c echo.Context, response *dto.ApiResponse, errs []error) error {
	response.Errors = errs

	statusCode := http.StatusBadRequest
	if perrors.HasNotFound(errs) {
		statusCode = http.StatusNotFound
	}

	respStatus := dto.ApiResponseStatus{
		Code: statusCode,
		Type: http.StatusText(statusCode),
	}
	response.Status = respStatus

	return c.JSON(statusCode, response)
}
