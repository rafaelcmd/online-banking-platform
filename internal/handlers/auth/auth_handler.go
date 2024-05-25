package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rafaelcmd/online-banking-platform/infrastructure/aws/cognito"
	user "github.com/rafaelcmd/online-banking-platform/internal/domain/user"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to create user")

	var req user.User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Error decoding request body: %v\n", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = cognito.CreateUser(req.UserName, req.Password, req.Email)
	if err != nil {
		log.Printf("Error creating user: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to initiate auth user")

	var req user.User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Error decoding request body: %v\n", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = cognito.InitiateAuthUser(req.UserName, req.Password)
	if err != nil {
		log.Printf("Error initiating auth user: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User authenticated successfully"))
}
