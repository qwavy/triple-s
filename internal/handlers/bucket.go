package handlers

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	errors2 "triple-s/internal/errors"
	"triple-s/internal/models"
	"triple-s/internal/services"
)

type BucketHandler struct {
	service *services.BucketService
}

func NewBucketHandler(bucketService *services.BucketService) *BucketHandler {
	return &BucketHandler{service: bucketService}
}

func (h *BucketHandler) Create(w http.ResponseWriter, r *http.Request) {
	var b models.Bucket
	b.Name = r.PathValue("BucketName")

	json.NewDecoder(r.Body).Decode(&b)
	createdBucket, err := h.service.CreateNewBucket(b.Name)

	if err != nil {
		if errors.Is(err, errors2.ErrBucketAlreadyExists) {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(map[string]string{"error": "This name already taken"})

			return
		}

		if errors.Is(err, errors2.ErrBucketInvalidName) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Name should be 3-63 characters, only lowercase letters, numbers, hyphens, and periods"})

			return
		}

		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(createdBucket)
}

func (h *BucketHandler) List(w http.ResponseWriter, r *http.Request) {
	response, _ := h.service.List()

	x, err := xml.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	// Write
	w.Write(x)
}
