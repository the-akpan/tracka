package controllers

import (
	"encoding/json"
	"net/http"
)

type apiResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func parseResponse(res http.ResponseWriter, req *http.Request, formStruct interface{}) (apiResponse, error) {
	var err error

	if formStruct != nil {
		err = json.NewDecoder(req.Body).Decode(formStruct)
	}

	return apiResponse{}, err
}

func writeResponseStatus(res http.ResponseWriter, status *int) {
	res.WriteHeader(*status)
}
