package api

import (
	"encoding/json"
	"github.com/pcavezzan/sample-apiserver/log"
	"net/http"
)

func writeJson(w http.ResponseWriter, r *http.Request, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		log.Error("An error occurs while marshalling to json", err)
		writeDefaultErrorResponse(w)
		return
	}
	w.Write(b)
}
