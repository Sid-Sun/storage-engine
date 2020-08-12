package delete

import (
	"encoding/json"
	"fmt"
	"github.com/sid-sun/notes-api/pkg/api/contract/delete"
	"github.com/sid-sun/notes-api/pkg/api/handlers"
	"github.com/sid-sun/notes-api/pkg/api/service"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

// Handler handles all delete requests
func Handler(logger *zap.Logger, svc service.Service) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		logger.Info("[Delete] [attempt]")

		if request.Body == nil {
			logger.Info(fmt.Sprintf("[%s] %s", api, "Request body is empty"))
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			logger.Error(fmt.Sprintf("[%s] [%s] %s", api, "ReadAll", err.Error()))
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		var req delete.Request
		err = json.Unmarshal(data, &req)
		if err != nil {
			logger.Error(fmt.Sprintf("[%s] [%s] %s", api, "Unmarshal", err.Error()))
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		if req.ID == "" || req.Pass == "" {
			logger.Info(fmt.Sprintf("[%s] %s", api, "ID or Pass empty"))
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		d, err := svc.Get(req.ID)
		if err != nil {
			logger.Info(fmt.Sprintf("[%s] [%s] %s", api, "Get", err.Error()))
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		if d.IsEmpty() {
			logger.Info(fmt.Sprintf("[%s] [%s]", api, "DataIsEmpty"))
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		_, err = handlers.DecryptAAD(d, req.Pass)
		if err != nil && err.Error() == handlers.IncorrectPassError {
			logger.Info(fmt.Sprintf("[%s] [%s] %s", api, "DecryptAAD", err.Error()))
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		if err != nil {
			logger.Error(fmt.Sprintf("[%s] [%s] %s", api, "DecryptAAD", err.Error()))
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = svc.Delete(req.ID)
		if err != nil {
			logger.Info(fmt.Sprintf("[%s] [%s] %s", api, "Delete", err.Error()))
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		res := delete.Response{
			ID:   req.ID,
		}
		data, err = json.Marshal(res)
		if err != nil {
			logger.Error(fmt.Sprintf("[%s] [%s] %s", api, "Marshal", err.Error()))
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = writer.Write(data)
		if err != nil {
			logger.Error(fmt.Sprintf("[%s] [%s] %s", api, "Write", err.Error()))
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		logger.Info("[Delete] [success]")
	}
}
