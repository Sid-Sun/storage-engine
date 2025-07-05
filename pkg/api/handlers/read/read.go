package read

import (
	"encoding/json"
	"fmt"
	"github.com/sid-sun/storage-engine/pkg/api/contract/read"
	"github.com/sid-sun/storage-engine/pkg/api/handlers"
	"github.com/sid-sun/storage-engine/pkg/api/service"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

func Handler(logger *zap.Logger, svc service.Service) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		logger.Info("[Read] [attempt]")

		if request.Body == nil {
			logger.Info(fmt.Sprintf("[%s] %s", api, "ReadRequest body is empty"))
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			logger.Error(fmt.Sprintf("[%s] [%s] %s", api, "ReadAll", err.Error()))
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		var req read.ReadRequest
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

		aad, err := handlers.DecryptAAD(d, req.Pass)
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

		note, err := handlers.Decrypt(d.Note, aad)
		if err != nil {
			logger.Error(fmt.Sprintf("[%s] [%s] %s", api, "Decrypt", err.Error()))
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		res := read.ReadResponse{
			ID:   req.ID,
			Note: note,
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

		logger.Info("[Read] [success]")
	}
}
