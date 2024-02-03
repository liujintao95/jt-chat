package ctxdata

import (
	"context"
	"encoding/json"
)

// CtxKeyJwtUserId get uid from ctx
var CtxKeyJwtUserId = "jwtUserId"

func GetUidFromCtx(ctx context.Context) string {
	var (
		uid     string
		jsonUid json.Number
	)
	jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(json.Number)
	if ok {
		uid = jsonUid.String()
	}
	return uid
}
