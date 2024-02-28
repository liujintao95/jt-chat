package message

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"jt-chat/apps/message/api/internal/logic/message"
	"jt-chat/apps/message/api/internal/svc"
	"jt-chat/apps/message/api/internal/types"
)

func GetNextMsgListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetNextMsgListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := message.NewGetNextMsgListLogic(r.Context(), svcCtx)
		resp, err := l.GetNextMsgList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}