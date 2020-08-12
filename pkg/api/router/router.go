package router

import (
	"github.com/gorilla/mux"
	"github.com/sid-sun/notes-api/pkg/api/handlers"
	"github.com/sid-sun/notes-api/pkg/api/handlers/create"
	"github.com/sid-sun/notes-api/pkg/api/handlers/ping"
	"github.com/sid-sun/notes-api/pkg/api/handlers/read"
	"github.com/sid-sun/notes-api/pkg/api/service"
	"go.uber.org/zap"
)

// New returns a new instance of the router
func New(logger *zap.Logger, svc service.Service) *mux.Router {
	myRouter := mux.NewRouter()

	myRouter.Handle("/", handlers.WithContentType(ping.Handler(logger))).Methods("GET")
	myRouter.Handle("/create", handlers.WithContentType(create.Handler(logger, svc))).Methods("POST")
	myRouter.Handle("/read", handlers.WithContentType(read.Handler(logger, svc))).Methods("GET")

	return myRouter
}
