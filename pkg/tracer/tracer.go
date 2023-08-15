package tracer

import (
	"context"

	"go.elastic.co/apm"
)

const (
	middlewareType = "app.internal.middleware"
	handlerType    = "app.internal.handler"
	usecaseType    = "app.internal.usecase"
	repoType       = "app.internal.repo"
	postgresType   = "db.postgresql.query"
	redisType      = "db.redis"
)

func StartSpanMiddleware(ctx context.Context, name string) (*apm.Span, context.Context) {
	return apm.StartSpan(ctx, name, middlewareType)
}

func StartSpanHandler(ctx context.Context, name string) (*apm.Span, context.Context) {
	return apm.StartSpan(ctx, name, handlerType)
}

func StartSpanUsecase(ctx context.Context, name string) (*apm.Span, context.Context) {
	return apm.StartSpan(ctx, name, usecaseType)
}

func StartSpanRepo(ctx context.Context, name string) (*apm.Span, context.Context) {
	return apm.StartSpan(ctx, name, repoType)
}

func StartSpanPostgres(ctx context.Context, name string) (*apm.Span, context.Context) {
	return apm.StartSpan(ctx, name, postgresType)
}

func StartSpanRedis(ctx context.Context, name string) (*apm.Span, context.Context) {
	return apm.StartSpan(ctx, name, redisType)
}
