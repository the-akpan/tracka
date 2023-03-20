package controllers

import (
	"encoding/json"
	"internal/schemas"
	"net/http"
)

func AuthLogin(res http.ResponseWriter, req *http.Request) {
	responseStatus := http.StatusBadRequest
	schema := schemas.Login{}

	response, err := parseResponse(res, req, &schema)

	defer json.NewEncoder(res).Encode(&response)
	defer writeResponseStatus(res, &responseStatus)

	if err != nil {
		response.Message = UNABLE_TO_DECODE
		return
	}

	if errors := schema.Validate(); len(errors) != 0 {
		response.Message = BAD_REQUEST
		response.Data = errors
		return
	}

	responseStatus = http.StatusOK
}

func AuthLogout(res http.ResponseWriter, req *http.Request) {}

func AuthForgotPassword(res http.ResponseWriter, req *http.Request) {}
