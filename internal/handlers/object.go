package handlers

import (
	"errors"
	"io"
	"net/http"
	errors2 "triple-s/internal/errors"
	"triple-s/internal/pkg"
	"triple-s/internal/services"
)

type ObjectHandler struct {
	service *services.ObjectService
}

func NewObjectHandler(objetService *services.ObjectService) *ObjectHandler {
	return &ObjectHandler{service: objetService}
}

func (h *ObjectHandler) Create(w http.ResponseWriter, r *http.Request) {
	bucketName := r.PathValue("BucketName")
	objectKey := r.PathValue("ObjectKey")

	content, err := io.ReadAll(r.Body)
	if err != nil {
		pkg.SendMessage(w, http.StatusBadRequest, "Error: Server")
	}

	err = h.service.Create(bucketName, objectKey, content)

	if err != nil {
		pkg.SendMessage(w, http.StatusInternalServerError, "Error: Server")
	}
}

func (h *ObjectHandler) Delete(w http.ResponseWriter, r *http.Request) {
	bucketName := r.PathValue("BucketName")
	objectKey := r.PathValue("ObjectKey")

	err := h.service.Delete(bucketName, objectKey)

	if err != nil {
		if errors.Is(err, errors2.ErrBucketNotFound) {
			pkg.SendMessage(w, http.StatusNotFound, "Error: Bucket not found")
		}

		if errors.Is(err, errors2.ErrObjNotFound) {
			pkg.SendMessage(w, http.StatusNotFound, "Error: Object not found")
		}
		pkg.SendMessage(w, http.StatusInternalServerError, "Error: Server")
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ObjectHandler) Get(w http.ResponseWriter, r *http.Request) {
	bucketName := r.PathValue("BucketName")
	objectKey := r.PathValue("ObjectKey")

	content, contentType, err := h.service.Get(bucketName, objectKey)

	if err != nil {
		if errors.Is(err, errors2.ErrBucketNotFound) {
			pkg.SendMessage(w, http.StatusNotFound, "Error: Bucket not found")
		}

		if errors.Is(err, errors2.ErrObjNotFound) {
			pkg.SendMessage(w, http.StatusNotFound, "Error: Object not found")
		}
	}

	w.Header().Set("Content-Type", contentType)
	w.Write(content)
}
