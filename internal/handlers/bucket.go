package handlers

import (
	"encoding/xml"
	"net/http"

	"triple-s/internal/pkg"
)

func (h *Handler) CreateBucket(w http.ResponseWriter, r *http.Request) {
	bucketName := r.PathValue("BucketName")

	_, err := h.service.CreateNewBucket(bucketName)
	if err != nil {
		pkg.SendError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ListBucket(w http.ResponseWriter, r *http.Request) {
	response, err := h.service.ListBucket()
	if err != nil {
		pkg.SendError(w, err)
		return
	}

	// ИСПРАВЛЕНО: Явный возврат XML ответа, как требует ТЗ
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)

	// Оборачиваем ответ в XML (не забудь добавить теги xml в models.BucketsResult)
	xmlBytes, err := xml.MarshalIndent(response, "", "  ")
	if err != nil {
		pkg.SendError(w, err)
		return
	}

	w.Write([]byte(xml.Header))
	w.Write(xmlBytes)
}

func (h *Handler) DeleteBucket(w http.ResponseWriter, r *http.Request) {
	bucketName := r.PathValue("BucketName")
	err := h.service.DeleteBucket(bucketName)
	if err != nil {
		pkg.SendError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
