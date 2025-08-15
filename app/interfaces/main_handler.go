package interfaces

import "github.com/labstack/echo/v4"

type MainHandlerInterface interface {
	HealthCheck() echo.HandlerFunc
}
