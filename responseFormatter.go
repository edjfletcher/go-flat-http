package go_flat_http

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseFormatAsJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)

	if err != nil {
		log.Fatal(err)
	}
}
