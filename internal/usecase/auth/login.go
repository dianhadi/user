package auth

import (
	"context"
	"database/sql"

	"github.com/dianhadi/golib/tracer"
	"github.com/dianhadi/user/internal/entity"
	"github.com/dianhadi/user/pkg/errors"
	"github.com/dianhadi/user/pkg/utils"
)

func (u Usecase) Login(ctx context.Context, user entity.User) (string, error) {
	span, ctx := tracer.StartSpanUsecase(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	metadata := make(map[string]interface{})

	metadata["username"] = user.Username

	userData, err := u.repoUser.GetUserByUsername(ctx, user.Username, false)
	if err != nil && err != sql.ErrNoRows {
		err := errors.New(errors.StatusInternalError, err)
		errors.AddMetadata(err, metadata)
		return "", err
	}

	// no data found
	if err == sql.ErrNoRows {
		err := errors.New(errors.StatusAuthorizationInvalid, err)
		errors.AddMetadata(err, metadata)
		return "", err
	}

	err = utils.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))
	if err != nil {
		err := errors.New(errors.StatusAuthorizationInvalid, err)
		errors.AddMetadata(err, metadata)
		return "", err
	}

	token := utils.GenerateToken(user.Username)

	err = u.repoAuth.SetSession(ctx, token, userData.ID)
	if err != nil {
		err := errors.New(errors.StatusInternalError, err)
		errors.AddMetadata(err, metadata)
		return "", err
	}

	return token, nil
}
