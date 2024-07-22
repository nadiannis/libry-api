package utils

import (
	"fmt"
	"net/http"

	"github.com/nadiannis/libry-api/internal/domain/response"
	"github.com/rs/zerolog/log"
)

func errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	res := response.ErrorResponse{
		Status: "error",
		Error:  message,
	}

	err := WriteJSON(w, status, res, nil)
	if err != nil {
		req := fmt.Sprint(r.Method, " ", r.URL.String())
		log.Error().Str("request", req).Err(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	req := fmt.Sprintf("%s %s %s", r.Proto, r.Method, r.URL.RequestURI())
	log.Error().Str("request", req).Msg(err.Error())

	message := "server encountered a problem"
	errorResponse(w, r, http.StatusInternalServerError, message)
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	req := fmt.Sprintf("%s %s %s", r.Proto, r.Method, r.URL.RequestURI())
	log.Error().Str("request", req).Msg(err.Error())

	errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func FailedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	req := fmt.Sprintf("%s %s %s", r.Proto, r.Method, r.URL.RequestURI())
	log.Error().Str("request", req).Msg("invalid request body")

	errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}
