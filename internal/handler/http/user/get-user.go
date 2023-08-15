package user

import (
	"net/http"

	"github.com/dianhadi/user/internal/handler/helper"
	"github.com/dianhadi/user/pkg/tracer"
	"github.com/dianhadi/user/pkg/utils"
	"github.com/go-chi/chi"
)

func (h Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	span, ctx := tracer.StartSpanHandler(r.Context(), utils.GetCurrentFunctionName())
	defer span.End()

	id := chi.URLParam(r, "id")

	user, err := h.usecaseUser.GetUserByID(ctx, id)
	if err != nil {
		helper.Write(w, ctx, err, nil)
		return
	}

	// cleanup response
	user.Password = ""

	helper.Write(w, ctx, nil, user)
}

func (h Handler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	span, ctx := tracer.StartSpanHandler(r.Context(), utils.GetCurrentFunctionName())
	defer span.End()

	username := chi.URLParam(r, "username")

	user, err := h.usecaseUser.GetUserByUsername(ctx, username)
	if err != nil {
		helper.Write(w, ctx, err, nil)
		return
	}

	// cleanup response
	user.Password = ""

	helper.Write(w, ctx, nil, user)
}
