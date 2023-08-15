package user

import (
	"encoding/json"
	"net/http"

	"github.com/dianhadi/user/internal/entity"
	"github.com/dianhadi/user/internal/handler/helper"
	"github.com/dianhadi/user/pkg/errors"
	"github.com/dianhadi/user/pkg/tracer"
	"github.com/dianhadi/user/pkg/utils"
)

func (h Handler) Register(w http.ResponseWriter, r *http.Request) {
	span, ctx := tracer.StartSpanHandler(r.Context(), utils.GetCurrentFunctionName())
	defer span.End()

	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		err := errors.New(errors.StatusRequestBodyInvalid, err)
		helper.Write(w, ctx, err, nil)
		return
	}

	err = h.usecaseUser.ValidateRegistration(ctx, user)
	if err != nil {
		helper.Write(w, ctx, err, nil)
		return
	}

	err = h.usecaseUser.Register(ctx, user)
	if err != nil {
		helper.Write(w, ctx, err, nil)
		return
	}

	// cleanup response
	user.Password = ""

	helper.Write(w, ctx, nil, user)
}
