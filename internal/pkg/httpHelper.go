package pkg

import (
	"encoding/json"
	"net/http"
)

func SendMessage(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": msg})
}
