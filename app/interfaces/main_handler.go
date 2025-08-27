package interfaces

import echo "github.com/labstack/echo/v4"

type MainHandlerInterface interface {
	HealthCheck() echo.HandlerFunc
}
