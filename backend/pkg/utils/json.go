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

// Here we're just returning the exact error from .Decode, but in a real life application
// We wouldn't show this to the user, we'd reformat it first
func RequestFromJSON(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(&v)
}
