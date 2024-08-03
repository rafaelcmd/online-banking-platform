package identity

import "errors"

type IdentityServiceInterface interface {
	CreateIdentityService() error
	CreateUser(username, password string, attributes map[string]string) (string, error)
	AuthenticateUser(username, password string) (string, error)
	DeleteUser(username string) error
}

func NewIdentityService(serviceName string) (IdentityServiceInterface, error) {
	switch serviceName {
	case "cognito":
		service := NewCognitoService()
		return service, nil
	case "adb2c":
		service := NewADB2CService()
		return service, nil
	default:
		return nil, errors.New("unsupported identity service")
	}
}
