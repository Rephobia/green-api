package response

import (
	"encoding/json"
	"net/http"
)

func Success(w http.ResponseWriter, data any) {
	sendJson(w, http.StatusOK, data)
}

func Error(w http.ResponseWriter, status int, msg any) {
	sendJson(w, status, map[string]any{
		"error": msg,
	})
}

func sendJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
