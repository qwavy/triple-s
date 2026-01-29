package main

import (
	"net/http"
	"triple-s/internal/handlers"
	"triple-s/internal/services"
	storage2 "triple-s/internal/storage"
)

func main() {
	bucketStorage := storage2.NewBucketStorage("./data/", "data/buckets.csv")
	bucketService := services.NewBucketService(bucketStorage)
	bucketHandler := handlers.NewBucketHandler(bucketService)

	mux := http.NewServeMux()
	mux.HandleFunc("PUT /{BucketName}", bucketHandler.Create)
	mux.HandleFunc("GET /", bucketHandler.List)
	mux.HandleFunc("DELETE /{BucketName}", bucketHandler.Delete)

	objectStorage := storage2.NewObjectStorage("./data")
	objectService := services.NewObjectService(objectStorage, bucketStorage)
	objectHandler := handlers.NewObjectHandler(objectService)
	//objectService := services.NewObjectService(objectStorage, bucketStorage)

	//mux.HandleFunc("PUT /{BucketName}/{ObjectKey}", )
	mux.HandleFunc("DELETE /{BucketName}/{ObjectKey}", objectHandler.Delete)

	http.ListenAndServe(":8080", mux)
}
