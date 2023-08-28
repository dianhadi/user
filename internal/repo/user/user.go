package user

import (
	"context"

	"github.com/dianhadi/golib/tracer"
	"github.com/dianhadi/user/internal/entity"
	"github.com/dianhadi/user/pkg/utils"
)

// Insert insert row with initial user data
func (r Repo) Insert(ctx context.Context, user entity.User) error {
	span, ctx := tracer.StartSpanRepo(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	return r.insertDB(ctx, user)
}

// GetUserByUsername getting user data by id
func (r Repo) GetUserByID(ctx context.Context, id string, cache bool) (entity.User, error) {
	span, ctx := tracer.StartSpanRepo(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	if cache {
		user, err := r.getUserByIDCache(ctx, id)
		if err == nil && user.ID != "" {
			return user, nil
		}
	}

	user, err := r.getUserByIDDB(ctx, id)
	if err != nil {
		return entity.User{}, err
	}

	if cache {
		r.setUserByIDCache(ctx, id, user)
	}

	return user, nil
}

// GetUserByUsername getting user data by username
func (r Repo) GetUserByUsername(ctx context.Context, username string, cache bool) (entity.User, error) {
	span, ctx := tracer.StartSpanRepo(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	if cache {
		user, err := r.getUserByUsernameCache(ctx, username)
		if err == nil && user.ID != "" {
			return user, nil
		}
	}

	user, err := r.getUserByUsernameDB(ctx, username)
	if err != nil {
		return entity.User{}, err
	}

	if cache {
		r.setUserByUsernameCache(ctx, username, user)
	}

	return user, nil
}

// GetUserByEmail getting user data by email
func (r Repo) GetUserByEmail(ctx context.Context, email string, cache bool) (entity.User, error) {
	span, ctx := tracer.StartSpanRepo(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	if cache {
		user, err := r.getUserByEmailCache(ctx, email)
		if err == nil && user.ID != "" {
			return user, nil
		}
	}

	user, err := r.getUserByEmailDB(ctx, email)
	if err != nil {
		return entity.User{}, err
	}

	if cache {
		r.setUserByEmailCache(ctx, email, user)
	}

	return user, nil
}

// ChangePassword to main user table
func (r Repo) ChangePassword(ctx context.Context, user entity.User) error {
	span, ctx := tracer.StartSpanRepo(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	return r.changePasswordDB(ctx, user)
}
