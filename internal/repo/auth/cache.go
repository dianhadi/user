package auth

import (
	"context"
	"fmt"

	"github.com/dianhadi/user/internal/constant"
	"github.com/dianhadi/user/pkg/tracer"
	"github.com/dianhadi/user/pkg/utils"
)

func (r Repo) GetSession(ctx context.Context, token string) (string, error) {
	span, ctx := tracer.StartSpanRedis(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	key := fmt.Sprintf(constant.PrefixSession, token)
	return r.cache.Get(ctx, key)
}

func (r Repo) SetSession(ctx context.Context, token, username string) error {
	span, ctx := tracer.StartSpanRedis(ctx, utils.GetCurrentFunctionName())
	defer span.End()

	key := fmt.Sprintf(constant.PrefixSession, token)
	return r.cache.SetEx(ctx, key, username, int64(constant.OneHour))
}
