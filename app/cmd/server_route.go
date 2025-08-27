package cmd

import (
	_ "blog-api/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

//

func (server *ApiServer) route() {
	// health check
	server.echo.GET("/api/healthy", server.mainHandler.HealthCheck())
	server.echo.GET("/api/swagger/*", echoSwagger.WrapHandler)

	// article
	articleGroup := server.echo.Group("/api/articles")
	articleGroup.GET("", server.articleHandler.GetList())
	articleGroup.GET("/:id", server.articleHandler.GetDetail())
	articleGroup.POST("", server.articleHandler.Create())
}
