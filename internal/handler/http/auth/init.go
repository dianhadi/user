package auth

import (
	"context"

	"github.com/dianhadi/user/internal/entity"
)

type usecaseAuth interface {
	Login(ctx context.Context, user entity.User) (string, error)
	Authenticate(ctx context.Context, token string) (entity.User, error)
}

type ResponseData struct {
	Token string `json:"token"`
}

type Handler struct {
	usecaseAuth usecaseAuth
}

func New(auth usecaseAuth) (*Handler, error) {
	return &Handler{
		usecaseAuth: auth,
	}, nil
}
