package api

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sid-sun/notes-api/cmd/config"
	"github.com/sid-sun/notes-api/pkg/api/router"
	"github.com/sid-sun/notes-api/pkg/api/service"
	"github.com/sid-sun/notes-api/pkg/api/store"
	"go.uber.org/zap"
)

// StartServer starts the api, inits all the requited submodules and routine for shutdown
func StartServer(cfg config.Config, logger *zap.Logger) {
	mc, svc := getService(cfg.DBConfig, logger)
	r := router.New(logger, svc)

	srv := &http.Server{Addr: cfg.App.Address(), Handler: r}

	logger.Info(fmt.Sprintf("[StartServer] Listening on %s", cfg.App.Address()))
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Error(fmt.Sprintf("[StartServer] [ListenAndServe]: %s", err.Error()))
			panic(err)
		}
	}()

	gracefulShutdown(srv, logger, mc)
}

func getService(dbc config.DBConfig, logger *zap.Logger) (*mongo.Client, service.Service) {
	mc := store.NewClient(dbc, logger)
	client, mcl, err := mc.GetCollection()
	if err != nil {
		// If initial connection or ping fails, panic
		panic(err)
	}
	cl := store.NewCollection(mcl, logger, dbc)
	st := store.NewStore(cl)
	return client, service.NewNotesService(st, logger)

}

func gracefulShutdown(srv *http.Server, logger *zap.Logger, mc *mongo.Client) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	logger.Info("Attempting GracefulShutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		if err := srv.Shutdown(ctx); err != nil {
			logger.Error(fmt.Sprintf("[GracefulShutdown] [Shutdown] [Server]: %s", err.Error()))
			panic(err)
		}
	}()

	go func() {
		if err := mc.Disconnect(ctx); err != nil {
			logger.Error(fmt.Sprintf("[GracefulShutdown] [Shutdown] [MongoClient]: %s", err.Error()))
			panic(err)
		}
	}()
}
