package interfaces

import echo "github.com/labstack/echo/v4"

type ArticleHandlerInterface interface {
	GetList() echo.HandlerFunc
	GetDetail() echo.HandlerFunc
	Create() echo.HandlerFunc
}
