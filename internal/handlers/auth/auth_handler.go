package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rafaelcmd/online-banking-platform/internal/application/auth"
	user "github.com/rafaelcmd/online-banking-platform/internal/domain/user"
)

type AuthHandler struct {
	authService *auth.AuthAppService
}

func NewAuthHandler(authService *auth.AuthAppService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to create user")

	var req user.User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Error decoding request body: %v\n", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	_, err = h.authService.RegisterUser(r.Context(), req)
	if err != nil {
		log.Printf("Error creating user: %v\n", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]string{"message": "User created successfully"}
	json.NewEncoder(w).Encode(response)
}

func (h *AuthHandler) SignInHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to initiate auth user")

	var req user.User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Error decoding request body: %v\n", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	token, err := h.authService.LoginUser(r.Context(), req)
	if err != nil {
		log.Printf("Error initiating auth user: %v\n", err)
		http.Error(w, "Failed to sign in user", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": "User authenticated successfully", "token": token}
	json.NewEncoder(w).Encode(response)
}
