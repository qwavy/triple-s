package handlers

import (
	"net/http"
	"triple-s/internal/services"
)

type ObjectHandler struct {
	service *services.ObjectService
}

func NewObjectHandler(objetService *services.ObjectService) *ObjectHandler {
	return &ObjectHandler{service: objetService}
}

//func (h *BucketHandler) Create(w http.ResponseWriter, r *http.Request) {
//
//}

func (h *ObjectHandler) Delete(w http.ResponseWriter, r *http.Request) {
	bucketName := r.PathValue("BucketName")
	objectKey := r.PathValue("ObjectKey")

	h.service.Delete(bucketName, objectKey)
}
