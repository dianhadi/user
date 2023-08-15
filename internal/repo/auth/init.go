package auth

import (
	"context"
)

type cache interface {
	Get(ctx context.Context, key string) (string, error)
	SetEx(ctx context.Context, key, value string, exp int64) error
}

type Repo struct {
	cache cache
}

func New(cache cache) (Repo, error) {
	return Repo{
		cache: cache,
	}, nil
}
