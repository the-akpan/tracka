package controllers

import (
	"encoding/json"
	"internal/logic"
	"internal/schemas"
	"log"
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

	user, err := logic.Login(&schema)
	if err != nil {
		// TODO: better error checking in case of db failure/internal error
		responseStatus = http.StatusUnauthorized
		response.Message = UNAUTHORIZED
		return
	}

	var encoded string
	if encoded, err = sc.Encode(cookieName, map[string]string{"email": user.Email}); err != nil {
		log.Printf("Could not set cookie %+v\n", err)
		responseStatus = http.StatusInternalServerError
		response.Message = SERVER_ERROR
		return
	}

	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    encoded,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
	}

	responseStatus = http.StatusOK
	response.Message = "Logged in succesfully"

	http.SetCookie(res, cookie)
}

func AuthLogout(res http.ResponseWriter, req *http.Request) {}

func AuthForgotPassword(res http.ResponseWriter, req *http.Request) {}
