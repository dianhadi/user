package helper

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"go.elastic.co/apm"
)

func Common(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tx := apm.TransactionFromContext(ctx)

		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}
		tx.Context.SetLabel("request-id", requestID)
		ctx = context.WithValue(ctx, "request-id", requestID)

		appVersion := r.Header.Get("X-App-Version")
		if appVersion != "" {
			tx.Context.SetLabel("app-version", appVersion)
		}

		appSource := r.Header.Get("X-App-Source")
		if appSource != "" {
			tx.Context.SetLabel("app-source", appSource)

		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
