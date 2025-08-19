package cmd

import (
	"blog-api/handler"
	"blog-api/repository"
	"blog-api/service"

	"github.com/sirupsen/logrus"
)

func (server *ApiServer) dependenciesInjection() {
	server.mainHandler = handler.NewMainHandler()

	server.dbRepository = repository.NewDbRepository(server.mysqlDsn)
	err := server.dbRepository.InitializeDB()
	if err != nil {
		logrus.Fatal(err)
	}

	// article
	server.articleRepository = repository.NewArticleRepository(server.dbRepository.GetDB())
	server.articleService = service.NewArticleService(server.articleRepository)
	server.articleHandler = handler.NewArticleHandler(server.articleService)
}
