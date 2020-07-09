package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseToJSON(w http.ResponseWriter, r *http.Request, response interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if response != nil {
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, "Encoding to JSON failed", status)
		}
	}
}

func ErrorToJSON(w http.ResponseWriter, r *http.Request, error error, status int) {
	ResponseToJSON(w, r, map[string]string{"error": error.Error()}, status)
}

func RequestFromJSON(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(&v)
}
