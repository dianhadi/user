package helper

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dianhadi/user/pkg/errors"
	"github.com/dianhadi/user/pkg/log"
)

type response struct {
	Status  int         `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Write(w http.ResponseWriter, ctx context.Context, err error, data interface{}) {
	if err == nil {
		err = errors.New(errors.StatusOK, nil)
	}

	errWrap := errors.CastToErrorWrapper(err)
	// in case error not nil, not wrapped
	if errWrap == nil {
		err = errors.New(errors.StatusInternalError, err)
		errWrap = errors.CastToErrorWrapper(err)
	}

	if errWrap.IsLoggable() {
		logField := log.Fields{
			"caller":     errWrap.GetErrorLine(),
			"request-id": ctx.Value("request-id"),
		}
		if metadata := errWrap.GetMetadata(); len(metadata) > 0 {
			logField["metadata"] = metadata
		}
		log.ErrorWithFields(errWrap.Error(), logField)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errWrap.GetHttpStatus())
	_ = json.NewEncoder(w).Encode(response{
		Status:  int(errWrap.GetCode()),
		Message: errWrap.GetUserMessage(),
		Data:    data,
	})
}
