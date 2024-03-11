package message

import (
	"context"
	"github.com/jinzhu/copier"
	"jt-chat/apps/message/rpc/message"

	"jt-chat/apps/message/api/internal/svc"
	"jt-chat/apps/message/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNextMsgListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNextMsgListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNextMsgListLogic {
	return &GetNextMsgListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNextMsgListLogic) GetNextMsgList(req *types.GetNextMsgListReq) (resp *types.GetNextMsgListResp, err error) {
	var (
		rpcOut *message.GetNextMsgListOut
	)
	rpcOut, err = l.svcCtx.MsgRpc.GetNextMsgList(l.ctx, &message.GetNextMsgListIn{
		MsgId:    req.MsgId,
		TargetId: req.TargetId,
		Size:     req.Size,
	})
	if err != nil {
		return nil, err
	}
	_ = copier.Copy(&resp, rpcOut)
	return
}
