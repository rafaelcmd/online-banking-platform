package identity

type ADB2CService struct{}

func NewADB2CService() *ADB2CService {
	return &ADB2CService{}
}

func (c *ADB2CService) CreateIdentityService() error {
	return nil
}

func (c *ADB2CService) CreateUser(username, password string, attributes map[string]string) (string, error) {
	return "", nil
}

func (c *ADB2CService) AuthenticateUser(username, password string) (string, error) {
	return "", nil
}

func (c *ADB2CService) DeleteUser(username string) error {
	return nil
}
