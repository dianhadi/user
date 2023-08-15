package user

import (
	"context"
	"database/sql"
)

type database interface {
	Connect() (*sql.DB, error)
}

type cache interface {
	Expire(ctx context.Context, key string, exp int64) error
	Get(ctx context.Context, key string) (string, error)
	HGetAll(ctx context.Context, key string) ([]interface{}, error)
	HSet(ctx context.Context, key string, data interface{}) error
	SetEx(ctx context.Context, key, value string, exp int64) error

	ScanStruct(source []interface{}, target interface{}) error
}

type Repo struct {
	db    database
	cache cache
}

func New(db database, cache cache) (Repo, error) {
	return Repo{
		db:    db,
		cache: cache,
	}, nil
}
