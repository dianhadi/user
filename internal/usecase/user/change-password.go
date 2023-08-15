package user

import (
	"context"
	"database/sql"

	"github.com/dianhadi/user/internal/entity"
	errors "github.com/dianhadi/user/pkg/errors"
	"github.com/dianhadi/user/pkg/tracer"
	"github.com/dianhadi/user/pkg/utils"
)

func (u Usecase) CheckPassword(ctx context.Context, user entity.User, newPassword string) (bool, error) {
	span, ctx := tracer.StartSpanUsecase(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	metadata := make(map[string]interface{})

	metadata["user-id"] = user.ID

	userData, err := u.repoUser.GetUserByID(ctx, user.ID, false)
	if err != nil && err != sql.ErrNoRows {
		err := errors.New(errors.StatusInternalError, err)
		errors.AddMetadata(err, metadata)
		return false, err
	}

	// no data found
	if err == sql.ErrNoRows {
		err := errors.New(errors.StatusAuthorizationInvalid, err)
		errors.AddMetadata(err, metadata)
		return false, err
	}

	err = utils.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))
	if err != nil {
		err := errors.NewWithMessage(errors.StatusAuthorizationInvalid, "Old Password is not matched")
		errors.AddMetadata(err, metadata)
		return false, err
	}

	err = utils.CompareHashAndPassword([]byte(userData.Password), []byte(newPassword))
	if err == nil {
		err := errors.NewWithMessage(errors.StatusAuthorizationInvalid, "Same password")
		errors.AddMetadata(err, metadata)
		return false, err
	}

	return true, nil
}

func (u Usecase) ChangePassword(ctx context.Context, user entity.User, newPassword string) error {
	span, ctx := tracer.StartSpanUsecase(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	metadata := make(map[string]interface{})

	metadata["user-id"] = user.ID

	user.Password, _ = utils.HashPassword(newPassword)
	err := u.repoUser.ChangePassword(ctx, user)

	if err != nil {
		err := errors.New(errors.StatusInternalError, err)
		errors.AddMetadata(err, metadata)
		return err
	}

	return nil

}
