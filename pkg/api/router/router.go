package router

import (
	"github.com/gorilla/mux"
	"github.com/sid-sun/storage-engine/pkg/api/handlers"
	"github.com/sid-sun/storage-engine/pkg/api/handlers/create"
	"github.com/sid-sun/storage-engine/pkg/api/handlers/delete"
	"github.com/sid-sun/storage-engine/pkg/api/handlers/ping"
	"github.com/sid-sun/storage-engine/pkg/api/handlers/read"
	"github.com/sid-sun/storage-engine/pkg/api/handlers/updatenote"
	"github.com/sid-sun/storage-engine/pkg/api/handlers/updatepass"
	"github.com/sid-sun/storage-engine/pkg/api/service"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
	"net/http"
	"path/filepath"
)

// New returns a new instance of the router
func New(logger *zap.Logger, svc service.Service) *mux.Router {
	myRouter := mux.NewRouter()

	// Serve swagger.json
	myRouter.HandleFunc("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		absPath, _ := filepath.Abs("./swagger.json")
		logger.Info("Serving swagger.json from", zap.String("path", absPath))
		http.ServeFile(w, r, absPath)
	}).Methods("GET")

	// Serve Swagger UI
	myRouter.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// swagger:route GET / ping pingEndpoint
	// Ping endpoint to check API health.
	// responses:
	//   200: pingResponse
	//   500: genericError
	myRouter.Handle("/", handlers.WithContentJSON(ping.Handler(logger))).Methods("GET")

	// swagger:route POST /create create createNoteEndpoint
	// Create a new note.
	// responses:
	//   200: createResponse
	//   400: genericError
	//   500: genericError
	myRouter.Handle("/create", handlers.WithContentJSON(create.Handler(logger, svc))).Methods("POST")

	// swagger:route GET /read read readNoteEndpoint
	// Read an existing note.
	// responses:
	//   200: readResponse
	//   400: genericError
	//   403: genericError
	//   404: genericError
	//   500: genericError
	myRouter.Handle("/read", handlers.WithContentJSON(read.Handler(logger, svc))).Methods("GET")

	// swagger:route PUT /update/note update updatenoteEndpoint
	// Update an existing note.
	// responses:
	//   200: updatenoteResponse
	//   400: genericError
	//   403: genericError
	//   404: genericError
	//   500: genericError
	myRouter.Handle("/update/note", handlers.WithContentJSON(updatenote.Handler(logger, svc))).Methods("PUT")

	// swagger:route PATCH /update/pass update updatepassEndpoint
	// Update an existing note's password.
	// responses:
	//   200: updatepassResponse
	//   400: genericError
	//   403: genericError
	//   404: genericError
	//   500: genericError
	myRouter.Handle("/update/pass", handlers.WithContentJSON(updatepass.Handler(logger, svc))).Methods("PATCH")

	// swagger:route DELETE /delete delete deleteNoteEndpoint
	// Delete an existing note.
	// responses:
	//   200: deleteResponse
	//   400: genericError
	//   403: genericError
	//   404: genericError
	//   500: genericError
	myRouter.Handle("/delete", handlers.WithContentJSON(delete.Handler(logger, svc))).Methods("DELETE")

	return myRouter
}
