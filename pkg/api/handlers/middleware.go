package handlers

import "net/http"

// WithContentJSON middleware sets writer's Content-Type header to application/json 
func WithContentJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(writer, request)
	})
}
