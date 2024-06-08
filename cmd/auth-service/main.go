package main

import (
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gorilla/mux"
	v1 "github.com/rafaelcmd/online-banking-platform/api/v1"
	"github.com/rafaelcmd/online-banking-platform/internal/application/auth"
	"github.com/rafaelcmd/online-banking-platform/internal/config"
	authHandler "github.com/rafaelcmd/online-banking-platform/internal/handlers/auth"
	"github.com/rafaelcmd/online-banking-platform/internal/infrastructure/aws/cognito"
	"github.com/rs/cors"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	sess, err := session.NewSession()
	if err != nil {
		log.Fatalf("Failed to create AWS session: %v", err)
	}

	cognitoClient := cognito.NewCognitoClient(sess)

	cognitoService := cognito.NewCognitoService(cognitoClient, cfg.UserPoolClientId)

	authAppService := auth.NewAuthAppService(cognitoService)

	authHandler := authHandler.NewAuthHandler(authAppService)

	r := mux.NewRouter()

	v1.RegisterAuthRoutes(r, authHandler)

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
