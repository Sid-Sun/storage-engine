package api

import (
	"context"
	"fmt"
	"github.com/sid-sun/notes-api/cmd/config"
	"github.com/sid-sun/notes-api/pkg/api/router"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// StartServer starts the api, inits all the requited submodules and routine for shutdown
func StartServer(cfg config.Config, logger *zap.Logger) {
	r := router.New(logger)

	srv := &http.Server{Addr: cfg.App.Address(), Handler: r}

	logger.Info(fmt.Sprintf("[StartServer] Listening on %s", cfg.App.Address()))
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Error(fmt.Sprintf("[StartServer] [ListenAndServe]: %s", err.Error()))
			panic(err)
		}
	}()

	gracefulShutdown(srv, logger)
}

func gracefulShutdown(srv *http.Server, logger *zap.Logger) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	logger.Info("Attempting GracefulShutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		if err := srv.Shutdown(ctx); err != nil {
			logger.Error(fmt.Sprintf("[GracefulShutdown] [Shutdown]: %s", err.Error()))
			panic(err)
		}
	}()
}
