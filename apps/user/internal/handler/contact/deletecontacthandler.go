package contact

import (
	"jt-chat/common/xerr"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"jt-chat/apps/user/internal/logic/contact"
	"jt-chat/apps/user/internal/svc"
	"jt-chat/apps/user/internal/types"
)

func DeleteContactHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteContactReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, xerr.ErrHandler(xerr.CustomErr(400, r.Context(), err)))
			return
		}

		l := contact.NewDeleteContactLogic(r.Context(), svcCtx)
		resp, err := l.DeleteContact(&req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, xerr.ErrHandler(err))
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}