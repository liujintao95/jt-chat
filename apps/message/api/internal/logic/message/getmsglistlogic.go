package message

import (
	"context"
	"github.com/jinzhu/copier"
	"jt-chat/apps/message/rpc/message"

	"jt-chat/apps/message/api/internal/svc"
	"jt-chat/apps/message/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMsgListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMsgListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMsgListLogic {
	return &GetMsgListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMsgListLogic) GetMsgList(req *types.GetMsgListReq) (resp *types.GetMsgListResp, err error) {
	var (
		rpcOut *message.GetMsgListOut
	)
	rpcOut, err = l.svcCtx.MsgRpc.GetMsgList(l.ctx, &message.GetMsgListIn{
		TargetId:    req.TargetId,
		ContentLike: req.ContentLike,
		Page:        req.Page,
		Size:        req.Size,
	})
	if err != nil {
		return nil, err
	}
	_ = copier.Copy(&resp, rpcOut)
	return
}
