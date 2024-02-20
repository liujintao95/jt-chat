package contact

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"jt-chat/apps/user/api/internal/logic/contact"
	"jt-chat/apps/user/api/internal/svc"
	"jt-chat/apps/user/api/internal/types"
)

func GetGroupContactApplicationListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetGroupContactApplicationListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := contact.NewGetGroupContactApplicationListLogic(r.Context(), svcCtx)
		resp, err := l.GetGroupContactApplicationList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
