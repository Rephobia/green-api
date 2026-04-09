package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Rephobia/green-api-test-task/internal/dto"
	"github.com/Rephobia/green-api-test-task/internal/response"
	"github.com/Rephobia/green-api-test-task/internal/service/greenapi"
)

func GetSettings(w http.ResponseWriter, r *http.Request, req dto.BaseRequest) {
	data, err := greenapi.New(req.IdInstance, req.APITokenInstance).GetSettings()

	checkResponseError(w, data, err)
}

func GetStateInstance(w http.ResponseWriter, r *http.Request, req dto.BaseRequest) {
	data, err := greenapi.New(req.IdInstance, req.APITokenInstance).GetStateInstance()
	checkResponseError(w, data, err)
}

func SendMessage(w http.ResponseWriter, r *http.Request, req dto.SendMessageRequest) {
	data, err := greenapi.New(req.IdInstance, req.APITokenInstance).SendMessage(
		req.ChatID,
		req.Message,
	)
	checkResponseError(w, data, err)
}

func SendFileByUrl(w http.ResponseWriter, r *http.Request, req dto.SendFileRequest) {
	data, err := greenapi.New(req.IdInstance, req.APITokenInstance).SendFileByUrl(
		req.ChatID,
		req.UrlFile,
		req.FileName,
		req.Caption,
	)
	checkResponseError(w, data, err)
}

func checkResponseError(w http.ResponseWriter, data []byte, err error) {
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())

		return
	}

	response.Success(w, json.RawMessage(data))
}
