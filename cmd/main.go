package main

import (
	"flag"
	"fmt"
	"net/http"
	"triple-s/internal/handlers"
	"triple-s/internal/middleware"
	storage2 "triple-s/internal/repository"
	"triple-s/internal/services"

	"github.com/rs/cors"
)

func main() {
	flag.Parse()

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
		fmt.Println(`Simple Storage Service.

**Usage:**
    triple-s [-port <N>] [-dir <S>]  
    triple-s --help

**Options:**
- --help     Show this screen.
- --port N   Port number
- --dir S    Path to the directory`)
		flag.PrintDefaults()
	}

	flag.Parse()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	corsHandler := c.Handler(mux)
	wrappedMux := middleware.Logger(corsHandler)
	http.ListenAndServe(":8080", wrappedMux)
}
