package handlers

import (
	"io"
	"net/http"

	"triple-s/internal/pkg"
)

func (h *Handler) CreateObject(w http.ResponseWriter, r *http.Request) {
	bucketName := r.PathValue("BucketName")
	objectKey := r.PathValue("ObjectKey")

	content, err := io.ReadAll(r.Body)
	if err != nil {
		pkg.SendError(w, err)
		return
	}

	err = h.service.CreateObject(bucketName, objectKey, content)
	if err != nil {
		pkg.SendError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteObject(w http.ResponseWriter, r *http.Request) {
	bucketName := r.PathValue("BucketName")
	objectKey := r.PathValue("ObjectKey")

	err := h.service.DeleteObject(bucketName, objectKey)
	if err != nil {
		pkg.SendError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) GetObject(w http.ResponseWriter, r *http.Request) {
	bucketName := r.PathValue("BucketName")
	objectKey := r.PathValue("ObjectKey")

	content, contentType, err := h.service.GetObject(bucketName, objectKey)
	if err != nil {
		pkg.SendError(w, err)
		return
	}

	w.Header().Set("Content-Type", contentType)
	w.Write(content)
}
