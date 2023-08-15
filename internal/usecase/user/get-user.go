package user

import (
	"context"
	"database/sql"

	"github.com/dianhadi/user/internal/entity"
	errors "github.com/dianhadi/user/pkg/errors"
	"github.com/dianhadi/user/pkg/tracer"
	"github.com/dianhadi/user/pkg/utils"
)

func (u Usecase) GetUserByID(ctx context.Context, id string) (entity.User, error) {
	span, ctx := tracer.StartSpanUsecase(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	metadata := make(map[string]interface{})

	metadata["user-id"] = id

	user, err := u.repoUser.GetUserByID(ctx, id, true)
	if err != nil && err != sql.ErrNoRows {
		err := errors.New(errors.StatusInternalError, err)
		errors.AddMetadata(err, metadata)
		return entity.User{}, err
	}

	// no data found
	if err == sql.ErrNoRows {
		err := errors.New(errors.StatusDataNotFound, err)
		errors.AddMetadata(err, metadata)
		return entity.User{}, err
	}

	return user, nil
}

func (u Usecase) GetUserByUsername(ctx context.Context, username string) (entity.User, error) {
	span, ctx := tracer.StartSpanUsecase(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	metadata := make(map[string]interface{})

	metadata["username"] = username

	user, err := u.repoUser.GetUserByUsername(ctx, username, true)

	if err != nil && err != sql.ErrNoRows {
		err := errors.New(errors.StatusInternalError, err)
		errors.AddMetadata(err, metadata)
		return entity.User{}, err
	}

	// no data found
	if err == sql.ErrNoRows {
		err := errors.New(errors.StatusDataNotFound, err)
		errors.AddMetadata(err, metadata)
		return entity.User{}, err
	}

	return user, nil
}
