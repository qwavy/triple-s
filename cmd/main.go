package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"triple-s/internal/handlers"
	"triple-s/internal/middleware"
	storage2 "triple-s/internal/repository"
	"triple-s/internal/services"
)

func main() {
	storage := storage2.NewStore("./data/", "buckets.csv")
	service := services.NewService(storage)
	handler := handlers.NewHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("PUT /{BucketName}", handler.CreateBucket)
	mux.HandleFunc("GET /", handler.ListBucket)
	mux.HandleFunc("DELETE /{BucketName}", handler.DeleteBucket)

	mux.HandleFunc("GET /{BucketName}/{ObjectKey}", handler.GetObject)
	mux.HandleFunc("PUT /{BucketName}/{ObjectKey}", handler.CreateObject)
	mux.HandleFunc("DELETE /{BucketName}/{ObjectKey}", handler.DeleteObject)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Использование программы:\n")
		fmt.Fprintf(os.Stderr, "  myprogram [опции]\n\nОпции:\n")
		flag.PrintDefaults() // Выводит стандартный список флагов
	}

	flag.Parse()

	wrappedMux := middleware.Logger(mux)
	http.ListenAndServe(":8080", wrappedMux)
}
