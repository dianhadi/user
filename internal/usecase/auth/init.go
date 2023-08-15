package auth

import (
	"context"

	"github.com/dianhadi/user/internal/entity"
)

type repoUser interface {
	GetUserByUsername(ctx context.Context, username string, cache bool) (entity.User, error)
}

type repoAuth interface {
	GetSession(ctx context.Context, token string) (string, error)
	SetSession(ctx context.Context, token, username string) error
}

type Usecase struct {
	repoUser repoUser
	repoAuth repoAuth
}

func New(user repoUser, auth repoAuth) (*Usecase, error) {
	return &Usecase{
		repoUser: user,
		repoAuth: auth,
	}, nil
}
