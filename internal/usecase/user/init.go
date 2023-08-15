package user

import (
	"context"

	"github.com/dianhadi/user/internal/entity"
)

type repoUser interface {
	Insert(ctx context.Context, user entity.User) error
	GetUserByID(ctx context.Context, id string, cache bool) (entity.User, error)
	GetUserByUsername(ctx context.Context, username string, cache bool) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string, cache bool) (entity.User, error)
	ChangePassword(ctx context.Context, user entity.User) error
}

type Usecase struct {
	repoUser repoUser
}

func New(user repoUser) (*Usecase, error) {
	return &Usecase{
		repoUser: user,
	}, nil
}
