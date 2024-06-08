package cognito

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/rafaelcmd/online-banking-platform/internal/domain/user"
	"github.com/rafaelcmd/online-banking-platform/pkg/auth"
)

func NewCognitoClient(sess *session.Session) *cognitoidentityprovider.CognitoIdentityProvider {
	return cognitoidentityprovider.New(sess)
}

type CognitoService struct {
	cognitoClient    *cognitoidentityprovider.CognitoIdentityProvider
	userPoolClientId string
}

func NewCognitoService(client *cognitoidentityprovider.CognitoIdentityProvider, clientId string) *CognitoService {
	return &CognitoService{
		cognitoClient:    client,
		userPoolClientId: clientId,
	}
}

func (c *CognitoService) SignUp(ctx context.Context, user user.User) (string, error) {
	signUpInput := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(c.userPoolClientId),
		Username: aws.String(user.UserName),
		Password: aws.String(user.Password),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(user.Email),
			},
		},
	}

	_, err := c.cognitoClient.SignUp(signUpInput)
	if err != nil {
		log.Printf("Error during SignUp API call: %v\n", err)
		return "", err
	}
	log.Println("User created successfully with AWS Cognito")

	return "User created successfully", nil
}

func (c *CognitoService) SignIn(ctx context.Context, user user.User) (string, error) {
	initiateAuthInput := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		ClientId: aws.String(c.userPoolClientId),
		AuthParameters: map[string]*string{
			"USERNAME": aws.String(user.UserName),
			"PASSWORD": aws.String(user.Password),
		},
	}

	res, err := c.cognitoClient.InitiateAuth(initiateAuthInput)
	if err != nil {
		log.Printf("Error during Initiate Auth User API call: %v\n", err)
		return "", err
	}
	log.Println("User authenticated successfully with AWS Cognito")

	tokenManager := auth.NewTokenManager()
	err = tokenManager.SaveToken(user.UserName, res.AuthenticationResult.AccessToken)
	if err != nil {
		log.Printf("Error saving token: %v\n", err)
		return "", err
	}
	log.Println("Token saved successfully")

	return *res.AuthenticationResult.AccessToken, nil
}
