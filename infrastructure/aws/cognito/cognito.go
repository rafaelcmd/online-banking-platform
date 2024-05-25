package cognito

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/rafaelcmd/online-banking-platform/pkg/auth"
)

var (
	sess             *session.Session
	cognitoClient    *cognitoidentityprovider.CognitoIdentityProvider
	userPoolClientId string
)

func init() {
	var err error
	sess, err = session.NewSession()
	if err != nil {
		log.Fatalf("Error creating AWS session: %v", err)
	}

	cognitoClient = NewCognitoClient(sess)
	userPoolClientId = os.Getenv("USER_POOL_CLIENT_ID")
}

func NewCognitoClient(sess *session.Session) *cognitoidentityprovider.CognitoIdentityProvider {
	return cognitoidentityprovider.New(sess)
}

func CreateUser(username, password, email string) error {
	signUpInput := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(userPoolClientId),
		Username: aws.String(username),
		Password: aws.String(password),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(email),
			},
		},
	}

	_, err := cognitoClient.SignUp(signUpInput)
	if err != nil {
		log.Printf("Error during SignUp API call: %v\n", err)
		return err
	}
	log.Println("User created successfully with AWS Cognito")

	return nil
}

func InitiateAuthUser(username, password string) error {
	initiateAuthInput := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		ClientId: aws.String(userPoolClientId),
		AuthParameters: map[string]*string{
			"USERNAME": aws.String(username),
			"PASSWORD": aws.String(password),
		},
	}

	res, err := cognitoClient.InitiateAuth(initiateAuthInput)
	if err != nil {
		log.Printf("Error during Initiate Auth User API call: %v\n", err)
		return err
	}
	log.Println("User authenticated successfully with AWS Cognito")

	tokenManager := auth.NewTokenManager()
	err = tokenManager.SaveToken(username, res.AuthenticationResult.AccessToken)
	if err != nil {
		log.Printf("Error saving token: %v\n", err)
		return err
	}
	log.Println("Token saved successfully")

	return nil
}
