package message

import (
	"jt-chat/common/xerr"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"jt-chat/apps/message/api/internal/logic/message"
	"jt-chat/apps/message/api/internal/svc"
	"jt-chat/apps/message/api/internal/types"
)

func GetPreviousMsgListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPreviousMsgListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, xerr.ErrHandler(xerr.CustomErr(400, r.Context(), err)))
			return
		}

		l := message.NewGetPreviousMsgListLogic(r.Context(), svcCtx)
		resp, err := l.GetPreviousMsgList(&req)
		if err != nil {

			httpx.OkJsonCtx(r.Context(), w, xerr.ErrHandler(err))
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
