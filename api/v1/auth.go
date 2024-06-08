package v1

import (
	"github.com/gorilla/mux"
	"github.com/rafaelcmd/online-banking-platform/internal/handlers/auth"
)

func RegisterAuthRoutes(r *mux.Router, authHandler *auth.AuthHandler) {
	r.HandleFunc("/api/v1/register", authHandler.CreateUserHandler).Methods("POST")
	r.HandleFunc("/api/v1/auth", authHandler.SignInHandler).Methods("POST")
}
