package cmd

import (
	"blog-api/interfaces"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

const API_SERVER_DEFAULT_PORT string = "8080"

type ApiServer struct {
	runtimeEnv string
	apiRunType string
	logLevel   string

	mysqlDsn string

	mainHandler  interfaces.MainHandlerInterface
	dbRepository interfaces.DbRepositoryInterface

	articleRepository interfaces.ArticleRepositoryInterface
	articleService    interfaces.ArticleServiceInterface
	articleHandler    interfaces.ArticleHandlerInterface

	echo *echo.Echo
}

func (server *ApiServer) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = API_SERVER_DEFAULT_PORT
	}

	server.Start(port)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	log.Infof("signal waiting:")
	s := <-sigCh
	log.Infof("signal received:%s", s.String())
	server.Teardown()
}

func (server *ApiServer) Start(port string) {
	errLoadEnv := server.loadEnv()
	if len(errLoadEnv) != 0 {
		for _, e := range errLoadEnv {
			log.Fatal(e)
		}
		return
	}

	server.echo = echo.New()
	server.setMiddleware()
	server.dependenciesInjection()
	server.route()

	go func() {
		if err := server.echo.Start(":" + port); err != nil && err != http.ErrServerClosed {
			server.echo.Logger.Fatal("shuting down the server")
		}
	}()

	log.Info("server running with go routine")
}

func (server *ApiServer) Teardown() {
	log.Info("server graceful shutdown")
}
