package interfaces

import "github.com/labstack/echo/v4"

type ArticleHandlerInterface interface {
	GetList() echo.HandlerFunc
}
