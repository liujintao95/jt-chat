package message

import (
	"context"
	"github.com/jinzhu/copier"
	"jt-chat/apps/message/rpc/message"

	"jt-chat/apps/message/api/internal/svc"
	"jt-chat/apps/message/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadFileLogic {
	return &DownloadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadFileLogic) DownloadFile(req *types.DownloadFileReq) (resp *types.DownloadFileResp, err error) {
	var (
		rpcOut *message.DownloadFileOut
	)
	rpcOut, err = l.svcCtx.MsgRpc.DownloadFile(l.ctx, &message.DownloadFileIn{
		MsgId:    req.MsgId,
		FilePath: req.FilePath,
	})
	if err != nil {
		return nil, err
	}
	_ = copier.Copy(&resp, rpcOut)
	return
}
