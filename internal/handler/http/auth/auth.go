package auth

import (
	"net/http"

	"github.com/dianhadi/user/internal/handler/helper"
	"github.com/dianhadi/user/pkg/errors"
	"github.com/dianhadi/user/pkg/tracer"
	"github.com/dianhadi/user/pkg/utils"
)

func (h Handler) Authenticate(w http.ResponseWriter, r *http.Request) {
	span, ctx := tracer.StartSpanHandler(r.Context(), utils.GetCurrentFunctionName())
	defer span.End()

	err := r.ParseForm()
	if err != nil {
		err := errors.New(errors.StatusRequestBodyInvalid, err)
		helper.Write(w, ctx, err, nil)
		return
	}

	// Get values from the form data
	token := r.FormValue("token")

	user, err := h.usecaseAuth.Authenticate(ctx, token)
	if err != nil {
		helper.Write(w, ctx, err, nil)
		return
	}

	helper.Write(w, ctx, nil, user)
}

func (h Handler) AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		span, ctx := tracer.StartSpanMiddleware(r.Context(), utils.GetCurrentFunctionName())
		defer span.End()

		auth := r.Header.Get("Authorization")
		if auth == "" || len(auth) < 8 {
			err := errors.NewWithMessage(errors.StatusAuthorizationInvalid, "Token is required")
			helper.Write(w, ctx, err, nil)
			return
		}
		if auth[:6] != "Bearer" {
			err := errors.NewWithMessage(errors.StatusAuthorizationInvalid, "Token is invalid")
			helper.Write(w, ctx, err, nil)
			return
		}

		token := auth[7:]

		user, err := h.usecaseAuth.Authenticate(ctx, token)
		if err != nil {
			helper.Write(w, ctx, err, nil)
			return
		}

		r.Header.Set("User-ID", user.ID)

		next.ServeHTTP(w, r)
	})
}