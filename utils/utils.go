package utils

import (
	"github.com/json-iterator/go"
	"net/http"
)

type FormError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func WriteErr(w http.ResponseWriter, err string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	PanicErr(jsoniter.NewEncoder(w).Encode(FormError{Message: err, StatusCode: statusCode}))
}

func WriteJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	PanicErr(jsoniter.NewEncoder(w).Encode(data))
}
