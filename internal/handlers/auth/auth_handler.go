package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

var svc *cognitoidentityprovider.CognitoIdentityProvider

func init() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Region: aws.String("us-east-1"),
		},
	}))

	svc = cognitoidentityprovider.New(sess)
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ClientId string `json:"clientId"`
	Email    string `json:"email"`
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to create user")

	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Error decoding request body: %v\n", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	log.Printf("Request payload: %+v\n", req)

	err = CreateUser(req.Username, req.Password, req.ClientId, req.Email)
	if err != nil {
		log.Printf("Error creating user: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

func CreateUser(username, password, clientId, email string) error {
	signUpInput := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(clientId),
		Username: aws.String(username),
		Password: aws.String(password),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(email),
			},
		},
	}

	_, err := svc.SignUp(signUpInput)
	if err != nil {
		log.Printf("Error during SignUp API call: %v\n", err)
		return err
	}

	log.Println("User created successfully with AWS Cognito")
	return nil
}
