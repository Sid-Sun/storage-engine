package updatenote

import (
	"encoding/json"
	"fmt"
	"github.com/sid-sun/storage-engine/pkg/api/contract/db"
	"github.com/sid-sun/storage-engine/pkg/api/contract/updatenote"
	"github.com/sid-sun/storage-engine/pkg/api/handlers"
	"github.com/sid-sun/storage-engine/pkg/api/service"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

// Handler handles all update note requests
func Handler(logger *zap.Logger, svc service.Service) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		logger.Info("[UpdateNote] [attempt]")

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

		var req updatenote.Request
		err = json.Unmarshal(data, &req)
		if err != nil {
			logger.Error(fmt.Sprintf("[%s] [%s] %s", api, "Unmarshal", err.Error()))
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		if req.ID == "" || req.Pass == "" || req.Note == "" {
			logger.Info(fmt.Sprintf("[%s] %s", api, "Essential data missing"))
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
			writer.WriteHeader(http.StatusForbidden)
			return
		}
		if err != nil {
			logger.Error(fmt.Sprintf("[%s] [%s] %s", api, "DecryptAAD", err.Error()))
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		if req.NewPass == "" {
			req.NewPass = req.Pass
		}

		aad, hash, newNote, err := handlers.Encrypt(req.Note, req.NewPass)
		if err != nil {
			logger.Error(fmt.Sprintf("[%s] [%s] %s", api, "Encrypt", err.Error()))
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		newData := db.NewDataInstance(aad, hash, newNote)
		err = svc.Update(req.ID, newData)
		if err != nil {
			logger.Error(fmt.Sprintf("[%s] [%s] %s", api, "Update", err.Error()))
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		res := updatenote.Response{ID: req.ID}
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

		logger.Info("[UpdateNote] [success]")
	}
}
