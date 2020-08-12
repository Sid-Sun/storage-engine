package ping

import (
	"encoding/json"
	"fmt"
	"github.com/sid-sun/notes-api/pkg/api/contract/ping"
	"go.uber.org/zap"
	"net/http"
)

// Handler handles all ping requests
func Handler(logger *zap.Logger) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		logger.Info("[Ping] [attempt]")

		res := ping.Response{
			Message: message,
		}

		data, err := json.Marshal(res)
		if err != nil {
			logger.Error(fmt.Sprintf("[%s] [%s] %s", api, "Marshal", err.Error()))
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = writer.Write(data)
		if err != nil {
			logger.Error(fmt.Sprintf("[%s] [%s] %s", api, "Write", err.Error()))
			return
		}

		logger.Info("[Ping] [success]")
	}
}
