package main

import (
	"net/http"
	"triple-s/internal/handlers"
	"triple-s/internal/services"
	storage2 "triple-s/internal/storage"
)

func main() {
	storage := storage2.NewBucketStorage("data/buckets.csv")

	service := services.NewBucketService(storage)

	handler := handlers.NewBucketHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("PUT /{BucketName}", handler.Create)
	mux.HandleFunc("GET /", handler.List)

	http.ListenAndServe(":8080", mux)
}
