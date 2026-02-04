package pkg

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"triple-s/internal/models"
)

func StatusFromError(err error) int {
	if errors.Is(err, models.ErrBucketNotFound) || errors.Is(err, models.ErrObjNotFound) {
		return http.StatusNotFound
	}

	if errors.Is(err, models.ErrBucketInvalidName) || errors.Is(err, models.ErrObjInvalidKey) {
		return http.StatusBadRequest
	}

	if errors.Is(err, models.ErrBucketAlreadyExists) {
		return http.StatusConflict
	}

	if errors.Is(err, models.ErrBucketNotEmpty) {
		return http.StatusConflict
	}

	return http.StatusInternalServerError
}

func SendError(w http.ResponseWriter, err error) {
	fmt.Println(err)
	status := StatusFromError(err)
	SendMessage(w, status, err)
}

func SendMessage(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(status)
	w.Write([]byte(xml.Header))
	xml.NewEncoder(w).Encode(value)
}
