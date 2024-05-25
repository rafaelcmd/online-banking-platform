package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/rafaelcmd/online-banking-platform/api/v1"
	"github.com/rafaelcmd/online-banking-platform/internal/config"
	"github.com/rs/cors"
)

func main() {
	cfg := config.LoadConfig()

	r := mux.NewRouter()

	v1.RegisterAuthRoutes(r)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	log.Printf("AuthService starting on port %s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, handler))
}
