package utils

import (
	"encoding/json"
	"net/http"
	"os"
)

func WriteJsonResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Add("Content-Type", "application/json")
	if os.Getenv("MODE") == "DEV" {
		w.Header().Add("Access-Control-Allow-Origin", os.Getenv("CLIENT_URL"))
		w.Header().Add("Access-Control-Allow-Credentials", "true")
	}
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func WriteJsonServerErr(w http.ResponseWriter, err error) {
	WriteJsonResponse(
		w,
		http.StatusInternalServerError,
		map[string]string{"error": err.Error()},
	)
}

func WriteJsonUnauthorized(w http.ResponseWriter, err error) {
	WriteJsonResponse(
		w,
		http.StatusUnauthorized,
		map[string]string{"error": err.Error()},
	)
}
