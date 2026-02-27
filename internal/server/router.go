package server

import (
	"fmt"
	"net/http"

	"triple-s/internal/handlers"
	"triple-s/internal/repository"
	"triple-s/internal/services"
)

func Router(dir string) *http.ServeMux {
	dir += "/"
	fmt.Println(dir)

	storage := repository.NewStore(dir, "buckets.csv")
	service := services.NewService(storage)
	handler := handlers.NewHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("PUT /{BucketName}", handler.CreateBucket)
	mux.HandleFunc("GET /", handler.ListBucket)
	mux.HandleFunc("DELETE /{BucketName}", handler.DeleteBucket)

	mux.HandleFunc("GET /{BucketName}/{ObjectKey}", handler.GetObject)
	mux.HandleFunc("PUT /{BucketName}/{ObjectKey}", handler.CreateObject)
	mux.HandleFunc("DELETE /{BucketName}/{ObjectKey}", handler.DeleteObject)

	return mux
}
