package pkg

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"

	"triple-s/internal/models"
)

type S3Error struct {
	XMLName xml.Name `xml:"Error"`
	Code    string   `xml:"Code"`
	Message string   `xml:"Message"`
}

func parseError(err error) (int, S3Error) {
	if errors.Is(err, models.ErrBucketNotFound) {
		return http.StatusNotFound, S3Error{Code: "NoSuchBucket", Message: "The specified bucket does not exist"}
	}
	if errors.Is(err, models.ErrObjNotFound) {
		return http.StatusNotFound, S3Error{Code: "NoSuchKey", Message: "The specified key does not exist."}
	}
	if errors.Is(err, models.ErrBucketInvalidName) || errors.Is(err, models.ErrObjInvalidKey) {
		return http.StatusBadRequest, S3Error{Code: "InvalidBucketName", Message: "The specified bucket is not valid."}
	}
	if errors.Is(err, models.ErrBucketAlreadyExists) {
		return http.StatusConflict, S3Error{Code: "BucketAlreadyExists", Message: "The requested bucket name is not available."}
	}
	if errors.Is(err, models.ErrBucketNotEmpty) {
		return http.StatusConflict, S3Error{Code: "BucketNotEmpty", Message: "The bucket you tried to delete is not empty."}
	}

	return http.StatusInternalServerError, S3Error{Code: "InternalError", Message: "We encountered an internal error. Please try again."}
}

func SendError(w http.ResponseWriter, err error) {
	fmt.Println("Server error:", err)

	status, s3err := parseError(err)
	SendMessage(w, status, s3err)
}

func SendMessage(w http.ResponseWriter, status int, value any) {
	if status == http.StatusNoContent || value == "" {
		w.WriteHeader(status)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(status)
	w.Write([]byte(xml.Header))

	if err := xml.NewEncoder(w).Encode(value); err != nil {
		fmt.Println("XML encoding error:", err)
	}
}
