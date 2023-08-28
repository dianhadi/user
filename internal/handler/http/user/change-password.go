package user

import (
	"net/http"

	"github.com/dianhadi/golib/tracer"
	"github.com/dianhadi/user/internal/entity"
	"github.com/dianhadi/user/internal/handler/helper"
	"github.com/dianhadi/user/pkg/errors"
	"github.com/dianhadi/user/pkg/utils"
)

func (h Handler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	span, ctx := tracer.StartSpanHandler(r.Context(), utils.GetCurrentFunctionName())
	defer span.End()

	err := r.ParseForm()
	if err != nil {
		err := errors.New(errors.StatusRequestBodyInvalid, err)
		helper.Write(w, ctx, err, nil)
		return
	}

	// Get values from the form data
	oldPassword := r.FormValue("old-password")
	newPassword := r.FormValue("new-password")

	userID := r.Header.Get("User-ID")

	user := entity.User{
		ID:       userID,
		Password: oldPassword,
	}

	valid, err := h.usecaseUser.CheckPassword(ctx, user, newPassword)
	if err != nil {
		helper.Write(w, ctx, err, nil)
		return
	}

	if !valid {
		helper.Write(w, ctx, err, nil)
		return
	}

	err = h.usecaseUser.ChangePassword(ctx, user, newPassword)
	if err != nil {
		helper.Write(w, ctx, err, nil)
		return
	}

	helper.Write(w, ctx, nil, nil)
}
