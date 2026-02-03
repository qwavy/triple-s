package handlers

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	errors2 "triple-s/internal/errors"
	"triple-s/internal/models"
	"triple-s/internal/pkg"
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
			pkg.SendMessage(w, http.StatusConflict, "This name already taken")

			return
		}

		if errors.Is(err, errors2.ErrBucketInvalidName) {
			pkg.SendMessage(w, http.StatusBadRequest, "Name should be 3-63 characters, only lowercase letters, numbers, hyphens, and periods")

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
	w.Write(x)
}

func (h *BucketHandler) Delete(w http.ResponseWriter, r *http.Request) {
	bucketName := r.PathValue("BucketName")
	err := h.service.Delete(bucketName)

	if err != nil {
		if errors.Is(err, errors2.ErrBucketNotFound) {
			pkg.SendMessage(w, http.StatusNotFound, "Bucket not found")
		}

		if errors.Is(err, errors2.ErrBucketNotEmpty) {
			pkg.SendMessage(w, http.StatusConflict, "Bucket is not empty")
		}
		fmt.Println(err)

		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

}
