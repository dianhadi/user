package user

import (
	"context"
	"fmt"
	"time"

	"github.com/dianhadi/golib/tracer"
	"github.com/dianhadi/user/internal/constant"
	"github.com/dianhadi/user/internal/entity"
	"github.com/dianhadi/user/pkg/utils"
)

type tempUser struct {
	ID         string
	Username   string
	Email      string
	Name       string
	Password   string
	Status     int8
	CreatedAt  string
	EnabledAt  string
	DisabledAt string
}

func (r Repo) getUserByIDCache(ctx context.Context, id string) (entity.User, error) {
	span, ctx := tracer.StartSpanRedis(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	key := fmt.Sprintf(constant.PrefixUserByID, id)

	values, err := r.cache.HGetAll(ctx, key)
	if err != nil {
		return entity.User{}, err
	}

	temp := tempUser{}
	r.cache.ScanStruct(values, &temp)

	user := mappingUser(temp)

	return user, nil
}

func (r Repo) getUserByUsernameCache(ctx context.Context, username string) (entity.User, error) {
	span, ctx := tracer.StartSpanRedis(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	key := fmt.Sprintf(constant.PrefixUserByUsername, username)

	values, err := r.cache.HGetAll(ctx, key)
	if err != nil {
		return entity.User{}, err
	}

	temp := tempUser{}
	r.cache.ScanStruct(values, &temp)

	user := mappingUser(temp)

	return user, nil
}

func (r Repo) getUserByEmailCache(ctx context.Context, email string) (entity.User, error) {
	span, ctx := tracer.StartSpanRedis(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	key := fmt.Sprintf(constant.PrefixUserByEmail, email)

	values, err := r.cache.HGetAll(ctx, key)

	if err != nil {
		return entity.User{}, err
	}

	temp := tempUser{}
	r.cache.ScanStruct(values, &temp)

	user := mappingUser(temp)

	return user, nil
}

func (r Repo) setUserByIDCache(ctx context.Context, id string, data entity.User) error {
	span, ctx := tracer.StartSpanRedis(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	key := fmt.Sprintf(constant.PrefixUserByID, id)
	err := r.cache.HSet(ctx, key, data)
	if err != nil {
		return err
	}

	return r.cache.Expire(ctx, key, int64(constant.OneMonth))
}

func (r Repo) setUserByUsernameCache(ctx context.Context, username string, data entity.User) error {
	span, ctx := tracer.StartSpanRedis(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	key := fmt.Sprintf(constant.PrefixUserByUsername, username)
	err := r.cache.HSet(ctx, key, data)
	if err != nil {
		return err
	}

	return r.cache.Expire(ctx, key, int64(constant.OneMonth))
}

func (r Repo) setUserByEmailCache(ctx context.Context, email string, data entity.User) error {
	span, ctx := tracer.StartSpanRedis(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	key := fmt.Sprintf(constant.PrefixUserByEmail, email)
	err := r.cache.HSet(ctx, key, data)
	if err != nil {
		return err
	}

	return r.cache.Expire(ctx, key, int64(constant.OneMonth))
}

func mappingUser(temp tempUser) entity.User {
	user := entity.User{
		ID:       temp.ID,
		Username: temp.Username,
		Email:    temp.Email,
		Name:     temp.Name,
		Password: temp.Password,
		Status:   temp.Status,
	}

	user.CreatedAt, _ = time.Parse(constant.TimeFormat, temp.CreatedAt)

	if temp.EnabledAt != "" {
		enabledAt, _ := time.Parse(constant.TimeFormat, temp.EnabledAt)
		user.EnabledAt = &enabledAt
	} else {
		user.EnabledAt = nil
	}
	if temp.DisabledAt != "" {
		disabledAt, _ := time.Parse(constant.TimeFormat, temp.EnabledAt)
		user.DisabledAt = &disabledAt
	} else {
		user.DisabledAt = nil
	}
	return user
}
