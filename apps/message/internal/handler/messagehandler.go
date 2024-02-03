package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"jt-chat/apps/message/internal/logic"
	"jt-chat/apps/message/internal/svc"
	"jt-chat/apps/message/internal/types"
)

func MessageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewMessageLogic(r.Context(), svcCtx)
		resp, err := l.Message(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
