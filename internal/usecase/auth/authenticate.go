package auth

import (
	"context"

	"github.com/dianhadi/golib/redis"
	"github.com/dianhadi/golib/tracer"
	"github.com/dianhadi/user/internal/entity"
	"github.com/dianhadi/user/pkg/errors"
	"github.com/dianhadi/user/pkg/utils"
)

func (u Usecase) Authenticate(ctx context.Context, token string) (entity.User, error) {
	span, ctx := tracer.StartSpanUsecase(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	userID, err := u.repoAuth.GetSession(ctx, token)
	if err != nil && err != redis.ErrNil {
		err := errors.New(errors.StatusInternalError, err)
		return entity.User{}, err
	}

	// no data found
	if err == redis.ErrNil {
		err := errors.New(errors.StatusAuthorizationInvalid, err)
		return entity.User{}, err
	}

	user := entity.User{
		ID: userID,
	}

	return user, nil
}
