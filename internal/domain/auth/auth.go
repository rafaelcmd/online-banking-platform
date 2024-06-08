package auth

import (
	"context"

	"github.com/rafaelcmd/online-banking-platform/internal/domain/user"
)

type AuthService interface {
	SignUp(ctx context.Context, user user.User) (string, error)
	SignIn(ctx context.Context, user user.User) (string, error)
}
