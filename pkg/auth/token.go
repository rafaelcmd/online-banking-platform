package auth

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

type TokenManager struct {
	ssmClient *ssm.SSM
}

func NewTokenManager() *TokenManager {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"),
	}))

	return &TokenManager{ssmClient: ssm.New(sess)}
}

func (tm *TokenManager) SaveToken(userId string, token *string) error {
	_, err := tm.ssmClient.PutParameter(&ssm.PutParameterInput{
		Name:      aws.String("/token/" + userId),
		Value:     token,
		Overwrite: aws.Bool(true),
		Type:      aws.String("SecureString"),
	})
	if err != nil {
		return err
	}

	return nil
}
