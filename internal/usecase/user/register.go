package user

import (
	"context"
	"database/sql"
	"time"

	"github.com/dianhadi/golib/tracer"
	"github.com/dianhadi/user/internal/entity"
	errors "github.com/dianhadi/user/pkg/errors"
	"github.com/dianhadi/user/pkg/utils"
	"github.com/google/uuid"
)

func (u Usecase) ValidateRegistration(ctx context.Context, user entity.User) error {
	span, ctx := tracer.StartSpanUsecase(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	metadata := make(map[string]interface{})

	metadata["username"] = user.Username
	metadata["email"] = user.Email

	if !utils.IsAlphaNumeric(user.Username) {
		err := errors.NewWithMessage(errors.StatusRequestBodyInvalid, "Username is not alphanumeric")
		errors.AddMetadata(err, metadata)
		return err
	}

	if !utils.IsValidEmail(user.Email) {
		err := errors.NewWithMessage(errors.StatusRequestBodyInvalid, "Email format is not valid")
		errors.AddMetadata(err, metadata)
		return err
	}

	_, err := u.repoUser.GetUserByUsername(ctx, user.Username, true)
	if err == nil {
		err := errors.NewWithMessage(errors.StatusRequestBodyInvalid, "Username is already taken")
		errors.AddMetadata(err, metadata)
		return err
	}

	if err != nil && err != sql.ErrNoRows {
		err := errors.New(errors.StatusInternalError, err)
		errors.AddMetadata(err, metadata)
		return err
	}

	_, err = u.repoUser.GetUserByEmail(ctx, user.Email, true)
	if err == nil {
		err := errors.NewWithMessage(errors.StatusRequestBodyInvalid, "Email is already registered")
		errors.AddMetadata(err, metadata)
		return err
	}

	if err != nil && err != sql.ErrNoRows {
		err := errors.New(errors.StatusInternalError, err)
		errors.AddMetadata(err, metadata)
		return err
	}

	return nil
}

func (u Usecase) Register(ctx context.Context, user entity.User) error {
	span, ctx := tracer.StartSpanUsecase(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	metadata := make(map[string]interface{})

	metadata["username"] = user.Username
	metadata["email"] = user.Email

	user.ID = uuid.New().String()
	user.Password, _ = utils.HashPassword(user.Password)
	user.CreatedAt = time.Now()
	user.Status = 0
	err := u.repoUser.Insert(ctx, user)

	if err != nil {
		err := errors.New(errors.StatusInternalError, err)
		errors.AddMetadata(err, metadata)
		return err
	}

	return nil
}
