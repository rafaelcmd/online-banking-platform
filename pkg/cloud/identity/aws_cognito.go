package identity

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/rafaelcmd/online-banking-platform/pkg/utils"
)

type CognitoService struct{}

func NewCognitoService() *CognitoService {
	return &CognitoService{}
}

func (c *CognitoService) CreateIdentityService() error {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-2")})
	if err != nil {
		return err
	}

	cfClient := cloudformation.New(sess)
	ssmClient := ssm.New(sess)

	res, err := ssmClient.GetParameter(&ssm.GetParameterInput{
		Name: aws.String("/obp/cognito/user-pool-template-url"),
	})
	if err != nil {
		return err
	}

	userPoolTemplate, err := utils.FetchTemplateURL(res.Parameter.Value)
	if err != nil {
		return err
	}

	userPoolStackName := "obp-user-pool-stack"
	_, err = cfClient.CreateStack(&cloudformation.CreateStackInput{
		StackName:    aws.String(userPoolStackName),
		TemplateBody: aws.String(userPoolTemplate),
		Capabilities: []*string{aws.String("CAPABILITY_NAMED_IAM")},
	})
	if err != nil {
		return err
	}

	res, err = ssmClient.GetParameter(&ssm.GetParameterInput{
		Name: aws.String("/obp/cognito/user-pool-client-template-url"),
	})
	if err != nil {
		return err
	}

	userPoolClientTemplate, err := utils.FetchTemplateURL(res.Parameter.Value)
	if err != nil {
		return err
	}

	userPoolClientStackName := "obp-user-pool-client-stack"
	_, err = cfClient.CreateStack(&cloudformation.CreateStackInput{
		StackName:    aws.String(userPoolClientStackName),
		TemplateBody: aws.String(userPoolClientTemplate),
		Capabilities: []*string{aws.String("CAPABILITY_NAMED_IAM")},
	})
	if err != nil {
		return err
	}

	err = cfClient.WaitUntilStackCreateComplete(&cloudformation.DescribeStacksInput{
		StackName: aws.String(userPoolClientStackName),
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *CognitoService) CreateUser(username, password string, attributes map[string]string) (string, error) {
	// Implementation using Cognito API
	return "", nil
}

func (c *CognitoService) AuthenticateUser(username, password string) (string, error) {
	// Implementation using Cognito API
	return "", nil
}

func (c *CognitoService) DeleteUser(username string) error {
	// Implementation using Cognito API
	return nil
}
