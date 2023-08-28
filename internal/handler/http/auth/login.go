package auth

import (
	"net/http"

	"github.com/dianhadi/golib/tracer"
	"github.com/dianhadi/user/internal/entity"
	"github.com/dianhadi/user/internal/handler/helper"
	"github.com/dianhadi/user/pkg/errors"
	"github.com/dianhadi/user/pkg/utils"
)

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	span, ctx := tracer.StartSpanHandler(r.Context(), utils.GetCurrentFunctionName())
	defer span.End()

	err := r.ParseForm()
	if err != nil {
		err := errors.New(errors.StatusRequestBodyInvalid, err)
		helper.Write(w, ctx, err, nil)
		return
	}

	// Get values from the form data
	user := entity.User{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	token, err := h.usecaseAuth.Login(ctx, user)
	if err != nil {
		helper.Write(w, ctx, err, nil)
		return
	}

	data := ResponseData{
		Token: token,
	}

	helper.Write(w, ctx, nil, data)
}
