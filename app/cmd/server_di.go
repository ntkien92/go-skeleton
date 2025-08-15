package cmd

import "blog-api/handler"

func (server *ApiServer) dependenciesInjection() {
	server.mainHandler = handler.NewMainHandler()

	server.articleHandler = handler.NewArticleHandler()
}
