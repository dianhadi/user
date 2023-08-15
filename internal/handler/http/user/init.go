package user

import (
	"context"

	"github.com/dianhadi/user/internal/entity"
)

type usecaseUser interface {
	Register(ctx context.Context, user entity.User) error
	ValidateRegistration(ctx context.Context, user entity.User) error
	CheckPassword(ctx context.Context, user entity.User, newPassword string) (bool, error)
	ChangePassword(ctx context.Context, user entity.User, newPassword string) error
	GetUserByID(ctx context.Context, id string) (entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (entity.User, error)
}

type Handler struct {
	usecaseUser usecaseUser
}

func New(user usecaseUser) (*Handler, error) {
	return &Handler{
		usecaseUser: user,
	}, nil
}
