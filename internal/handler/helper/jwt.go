package helper

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dianhadi/user/internal/config"
	"github.com/dianhadi/user/pkg/errors"
	"github.com/dianhadi/user/pkg/tracer"
	"github.com/dianhadi/user/pkg/utils"
	"go.elastic.co/apm"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		span, ctx := tracer.StartSpanMiddleware(r.Context(), utils.GetCurrentFunctionName())
		defer span.End()
		tx := apm.TransactionFromContext(ctx)

		auth := r.Header.Get("Authorization")
		if auth == "" {
			err := errors.NewWithMessage(errors.StatusAuthorizationInvalid, "Token is required")
			Write(w, ctx, err, nil)
			return
		}
		tokenString := utils.GetTokenFromHeader(r)
		if tokenString == "" {
			err := errors.NewWithMessage(errors.StatusAuthorizationInvalid, "Token is invalid")
			Write(w, ctx, err, nil)
			return
		}

		token, err := validateToken(tokenString)
		if err != nil {
			err := errors.NewWithMessage(errors.StatusAuthorizationInvalid, "Token is invalid")
			Write(w, ctx, err, nil)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		service := claims["service"].(string)

		tx.Context.SetLabel("service", service)
		ctx = context.WithValue(ctx, "service", service)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func validateToken(tokenString string) (*jwt.Token, error) {
	publicKey := config.PublicKey

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return token, nil
}
