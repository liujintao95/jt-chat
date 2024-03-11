package message

import (
	"bytes"
	"fmt"
	"jt-chat/common/xerr"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/rest/httpx"
	"jt-chat/apps/message/api/internal/logic/message"
	"jt-chat/apps/message/api/internal/svc"
	"jt-chat/apps/message/api/internal/types"
)

func DownloadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadFileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, xerr.ErrHandler(xerr.CustomErr(400, r.Context(), err)))
			return
		}

		l := message.NewDownloadFileLogic(r.Context(), svcCtx)
		resp, err := l.DownloadFile(&req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, xerr.ErrHandler(err))
		} else {
			w.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, resp.Name))
			w.Header().Add("Content-Type", "application/octet-stream")
			w.Header().Add("Content-Length", fmt.Sprintf("%d", resp.Size))
			http.ServeContent(w, r, resp.Name, time.Now(), bytes.NewReader(resp.Data))
		}
	}
}
