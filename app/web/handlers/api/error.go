package api

import (
	"encoding/json"
	"github.com/pcavezzan/sample-apiserver/log"
	"net/http"
)


type ErrorResponse struct {
	Msg string	`json:"msg"`
	Code *int	`json:"code,omitempty"`
}

const (
	defaultMsg =  "An error occurs while processing your request."
)

func defaultErrorResponse() ErrorResponse {
	return ErrorResponse{Msg:defaultMsg}
}

func writeDefaultErrorResponse(w http.ResponseWriter) {
	 WriteError(w, defaultErrorResponse())
}

func WriteError(w http.ResponseWriter, e ErrorResponse) {
	w.WriteHeader(http.StatusInternalServerError)
	r, err := json.Marshal(e)
	if err != nil {
		log.Error("Could not marshalling error response in JSON.", err)
		e, _ := json.Marshal(defaultErrorResponse())
		r = []byte(e)
	}
	w.Write(r)
}