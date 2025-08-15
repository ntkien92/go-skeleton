package cmd

func (server *ApiServer) route() {
	// health check
	server.echo.GET("/api/healthy", server.mainHandler.HealthCheck())

	// article
	articleGroup := server.echo.Group("/api/articles")
	articleGroup.GET("", server.articleHandler.GetList())
}
