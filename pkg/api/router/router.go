package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

// New returns a new instance of the router
func New(logger *zap.Logger) *mux.Router {
	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		logger.Info("[Ping] [attempt]")
		_, err := writer.Write([]byte("Well, Hello!"))

		if err != nil {
			logger.Error(fmt.Sprintf("[%s], [%s] %s", "Router", "New", err.Error()))
		}
		logger.Info("[Ping] [success]")
	}).Methods("GET")

	return myRouter
}
