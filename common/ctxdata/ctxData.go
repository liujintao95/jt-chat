package ctxdata

import (
	"context"
)

// CtxKeyJwtUserId get uid from ctx
var CtxKeyJwtUserId = "jwtUserId"

func GetUidFromCtx(ctx context.Context) string {
	return ctx.Value(CtxKeyJwtUserId).(string)
}

func GenerateCtx(ctx context.Context, uid string) context.Context {
	return context.WithValue(ctx, CtxKeyJwtUserId, uid)
}
