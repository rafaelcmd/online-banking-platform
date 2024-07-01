package main

import (
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/gorilla/mux"
	v1 "github.com/rafaelcmd/online-banking-platform/api/v1"
	"github.com/rafaelcmd/online-banking-platform/internal/application/auth"
	authHandler "github.com/rafaelcmd/online-banking-platform/internal/handlers/auth"
	"github.com/rafaelcmd/online-banking-platform/internal/infrastructure/aws/cognito"
	"github.com/rs/cors"
)

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"),
	}))
	svc := ssm.New(sess)

	cognitoClient := cognito.NewCognitoClient(sess)

	userPoolClientId, err := getParameter(svc, "USER_POOL_CLIENT_ID")
	if err != nil {
		log.Fatalf("Failed to get parameter: %v", err)
	}

	cognitoService := cognito.NewCognitoService(cognitoClient, userPoolClientId)

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

	log.Printf("AuthService started")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func getParameter(svc *ssm.SSM, name string) (string, error) {
	param, err := svc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(name),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return "", err
	}
	return *param.Parameter.Value, nil
}
