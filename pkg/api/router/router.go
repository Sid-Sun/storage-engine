package router

import (
	"github.com/gorilla/mux"
	"github.com/sid-sun/notes-api/pkg/api/handlers"
	"github.com/sid-sun/notes-api/pkg/api/handlers/create"
	"github.com/sid-sun/notes-api/pkg/api/handlers/delete"
	"github.com/sid-sun/notes-api/pkg/api/handlers/ping"
	"github.com/sid-sun/notes-api/pkg/api/handlers/read"
	"github.com/sid-sun/notes-api/pkg/api/handlers/updatenote"
	"github.com/sid-sun/notes-api/pkg/api/service"
	"go.uber.org/zap"
)

// New returns a new instance of the router
func New(logger *zap.Logger, svc service.Service) *mux.Router {
	myRouter := mux.NewRouter()

	myRouter.Handle("/", handlers.WithContentJSON(ping.Handler(logger))).Methods("GET")
	myRouter.Handle("/create", handlers.WithContentJSON(create.Handler(logger, svc))).Methods("POST")
	myRouter.Handle("/read", handlers.WithContentJSON(read.Handler(logger, svc))).Methods("GET")
	myRouter.Handle("/update/note", handlers.WithContentJSON(updatenote.Handler(logger, svc))).Methods("PUT")
	myRouter.Handle("/delete", handlers.WithContentJSON(delete.Handler(logger, svc))).Methods("DELETE")

	return myRouter
}
