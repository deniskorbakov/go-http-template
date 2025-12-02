package handler

import (
	"encoding/json"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "Hello World",
		"status":  "success",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
