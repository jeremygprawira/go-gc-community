package app

import (
	"context"
	"errors"
	"go-gc-community/internal/config"
	handler "go-gc-community/internal/delivery/http"
	"go-gc-community/internal/repository"
	"go-gc-community/internal/server"
	"go-gc-community/internal/usecase"
	"go-gc-community/pkg/database/mysql"
	"go-gc-community/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	config, err := config.Init()
	if err != nil {
		logger.Error(err)
		return
	}

	// Database
	sql, err := mysql.Connect(config.Sql.User, config.Sql.Password, config.Sql.Host, config.Sql.Name, config.Sql.Charset)
	if err != nil {
		logger.Error(err)
	}

	repository := repository.NewRepositories(sql)
	usecase := usecase.NewUsecases(usecase.Dependencies{
		Repository: repository,
	})
	handler := handler.NewHandler(usecase)
	
	srv := server.NewServer(config, handler.Init(config))

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()


	logger.Info("Server Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}
}