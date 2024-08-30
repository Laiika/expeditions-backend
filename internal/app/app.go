package app

import (
	"context"
	"db_cp_6/config"
	v1 "db_cp_6/internal/controller/http/v1"
	"db_cp_6/internal/httpserver"
	"db_cp_6/internal/repo"
	"db_cp_6/internal/service"
	"db_cp_6/pkg/logger"
	"db_cp_6/pkg/postgres"
	_ "db_cp_6/swagger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config, log *logger.Logger) {
	log.Info("initializing postgres")
	member, err := postgres.NewClient(context.Background(), 3, &cfg.Member)
	if err != nil {
		log.Fatal(err)
	}
	defer member.Close()

	leader, err := postgres.NewClient(context.Background(), 3, &cfg.Leader)
	if err != nil {
		log.Fatal(err)
	}
	defer leader.Close()

	admin, err := postgres.NewClient(context.Background(), 3, &cfg.Admin)
	if err != nil {
		log.Fatal(err)
	}
	defer admin.Close()
	log.Info("connected to db")

	log.Info("initializing repositories")
	repos := repo.NewRepositories()

	log.Info("initializing services")
	services := service.NewServices(repos, admin, leader, member)

	log.Info("initializing handlers and routes")
	handler := gin.Default()
	v1.NewRouter(handler, services, log)

	log.Info("starting http server")
	address := fmt.Sprintf("%s:%s", cfg.HTTPServer.Host, cfg.HTTPServer.Port)
	httpServer := httpserver.New(handler, address)
	go func() {
		err = httpServer.Start()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	err = httpServer.Stop()
	if err != nil {
		panic(errors.Wrap(err, "Httpserver shutdown"))
	}
	log.Debug("Httpserver exited")
}
