package create

import (
	"encoding/json"
	"fmt"
	"github.com/sid-sun/notes-api/pkg/api/contract/create"
	"github.com/sid-sun/notes-api/pkg/api/contract/db"
	"github.com/sid-sun/notes-api/pkg/api/handlers"
	"github.com/sid-sun/notes-api/pkg/api/service"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

// Handler handles all create requests
func Handler(logger *zap.Logger, svc service.Service) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		logger.Info("[Create] [attempt]")

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

		var req create.Request
		err = json.Unmarshal(data, &req)
		if err != nil {
			logger.Error(fmt.Sprintf("[%s] [%s] %s", api, "Unmarshal", err.Error()))
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		if req.Pass == "" || req.Note == "" {
			logger.Info(fmt.Sprintf("[%s] %s", api, "Pass or Note empty"))
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		exists, err := svc.Exists(req.ID)
		if err != nil {
			logger.Error(fmt.Sprintf("[%s] [%s] %s", api, "Exists", err.Error()))
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		if req.ID == "" || exists {
			req.ID = handlers.RandString(8)
		}

		aad, hash, note, err := handlers.Encrypt(req.Note, req.Pass)
		if err != nil {
			logger.Error(fmt.Sprintf("[%s] [%s] %s", api, "Encrypt", err.Error()))
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = svc.Create(req.ID, db.NewDataInstance(aad, hash, note))
		if err != nil {
			logger.Error(fmt.Sprintf("[%s] [%s] %s", api, "Create", err.Error()))
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		res := create.Response{ID: req.ID}
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

		logger.Info("[Create] [success]")
	}
}
