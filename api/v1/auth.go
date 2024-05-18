package v1

import (
	"github.com/gorilla/mux"
	"github.com/rafaelcmd/online-banking-platform/internal/handlers/auth"
)

func RegisterAuthRoutes(r *mux.Router) {
	r.HandleFunc("/api/v1/register", auth.CreateUserHandler).Methods("POST")
}
