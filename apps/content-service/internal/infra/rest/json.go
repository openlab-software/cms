package rest

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func writeErr(w http.ResponseWriter, err error) error {
	return writeJSON(w, http.StatusInternalServerError, map[string]any{"err": err})
}
