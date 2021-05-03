package http_utils

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	var err error
	if payload == nil {
		_, err = w.Write([]byte(""))
	} else {
		var body []byte
		body, err = json.Marshal(payload)
		if err != nil {
			panic(err)
		}
		_, err = w.Write(body)
	}
	w.WriteHeader(statusCode)
	if err != nil {
		panic(err)
	}
}
