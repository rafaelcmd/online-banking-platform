package auth

import (
	"context"
	"errors"

	"github.com/rafaelcmd/online-banking-platform/internal/domain/auth"
	"github.com/rafaelcmd/online-banking-platform/internal/domain/user"
)

type AuthAppService struct {
	authService auth.AuthService
}

func NewAuthAppService(authService auth.AuthService) *AuthAppService {
	return &AuthAppService{
		authService: authService,
	}
}

func (a *AuthAppService) RegisterUser(ctx context.Context, user user.User) (string, error) {
	if user.Email == "" || user.UserName == "" || user.Password == "" {
		return "", errors.New("invalid user data")
	}
	return a.authService.SignUp(ctx, user)
}

func (a *AuthAppService) LoginUser(ctx context.Context, user user.User) (string, error) {
	if user.Email == "" || user.UserName == "" || user.Password == "" {
		return "", errors.New("invalid user data")
	}
	return a.authService.SignIn(ctx, user)
}
