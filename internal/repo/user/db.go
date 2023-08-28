package user

import (
	"context"
	"time"

	"github.com/dianhadi/golib/tracer"
	"github.com/dianhadi/user/internal/entity"
	"github.com/dianhadi/user/pkg/utils"
)

func (r Repo) insertDB(ctx context.Context, user entity.User) error {
	span, ctx := tracer.StartSpanPostgres(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	db, err := r.db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO users (id, username, email, name, password, status, created_at) values($1,$2,$3,$4,$5,$6,$7)",
		user.ID,
		user.Username,
		user.Email,
		user.Name,
		user.Password,
		user.Status,
		time.Now())
	return err
}

func (r Repo) getUserByIDDB(ctx context.Context, id string) (entity.User, error) {
	span, ctx := tracer.StartSpanPostgres(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	db, err := r.db.Connect()
	if err != nil {
		return entity.User{}, err
	}
	defer db.Close()

	var user entity.User
	err = db.QueryRow("SELECT id, username, email, name, password, status, created_at, enabled_at, disabled_at FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Username, &user.Email, &user.Name, &user.Password, &user.Status, &user.CreatedAt, &user.EnabledAt, &user.DisabledAt)

	return user, err
}

func (r Repo) getUserByUsernameDB(ctx context.Context, username string) (entity.User, error) {
	span, ctx := tracer.StartSpanPostgres(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	db, err := r.db.Connect()
	if err != nil {
		return entity.User{}, err
	}
	defer db.Close()

	var user entity.User
	err = db.QueryRow("SELECT id, username, email, name, password, status, created_at, enabled_at, disabled_at FROM users WHERE username = $1", username).
		Scan(&user.ID, &user.Username, &user.Email, &user.Name, &user.Password, &user.Status, &user.CreatedAt, &user.EnabledAt, &user.DisabledAt)

	return user, err
}

func (r Repo) getUserByEmailDB(ctx context.Context, email string) (entity.User, error) {
	span, ctx := tracer.StartSpanPostgres(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	db, err := r.db.Connect()
	if err != nil {
		return entity.User{}, err
	}
	defer db.Close()

	var user entity.User
	err = db.QueryRow("SELECT id, username, email, name, password, status, created_at, enabled_at, disabled_at FROM users WHERE email = $1", email).
		Scan(&user.ID, &user.Username, &user.Email, &user.Name, &user.Password, &user.Status, &user.CreatedAt, &user.EnabledAt, &user.DisabledAt)

	return user, err
}

func (r Repo) changePasswordDB(ctx context.Context, user entity.User) error {
	span, ctx := tracer.StartSpanPostgres(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	db, err := r.db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE users SET password = $1 WHERE id = $2",
		user.Password,
		user.ID)
	return err
}
