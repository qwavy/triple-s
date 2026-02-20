package handlers

import (
	"encoding/json"
	"net/http"
	"triple-s/internal/models"
	"triple-s/internal/pkg"
)

func (h *Handler) CreateBucket(w http.ResponseWriter, r *http.Request) {
	var b models.Bucket
	b.Name = r.PathValue("BucketName")

	json.NewDecoder(r.Body).Decode(&b)
	createdBucket, err := h.service.CreateNewBucket(b.Name)

	if err != nil {
		pkg.SendError(w, err)
		return
	}

	pkg.SendMessage(w, http.StatusOK, createdBucket)
}

func (h *Handler) ListBucket(w http.ResponseWriter, r *http.Request) {
	response, err := h.service.ListBucket()

	if err != nil {
		pkg.SendError(w, err)
		return
	}

	if err != nil {
		pkg.SendError(w, err)
		return
	}
	pkg.SendMessage(w, http.StatusOK, response)
}

func (h *Handler) DeleteBucket(w http.ResponseWriter, r *http.Request) {
	bucketName := r.PathValue("BucketName")
	err := h.service.DeleteBucket(bucketName)

	if err != nil {
		pkg.SendError(w, err)
	}

	pkg.SendMessage(w, http.StatusNoContent, "")
}
