package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/rafaelcmd/online-banking-platform/api/v1"
	"github.com/rafaelcmd/online-banking-platform/internal/config"
)

func main() {
	cfg := config.LoadConfig()
	r := mux.NewRouter()

	v1.RegisterAuthRoutes(r)

	log.Printf("AuthService starting on port %s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
