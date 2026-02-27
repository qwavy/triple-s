package server

import (
	"fmt"
	"log"
	"net/http"

	"triple-s/internal/config"
)

func RunServer(cfg config.Config) error {
	router := Router(cfg.Dir)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: router,
	}

	fmt.Printf("Server started on http://localhost:%d\n", cfg.Port)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
	return nil
}
