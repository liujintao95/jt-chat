package message

import (
	"context"
	"github.com/jinzhu/copier"
	"jt-chat/apps/message/rpc/message"

	"jt-chat/apps/message/api/internal/svc"
	"jt-chat/apps/message/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPreviousMsgListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPreviousMsgListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPreviousMsgListLogic {
	return &GetPreviousMsgListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPreviousMsgListLogic) GetPreviousMsgList(req *types.GetPreviousMsgListReq) (resp *types.GetPreviousMsgListResp, err error) {
	var (
		rpcOut *message.GetPreviousMsgListOut
	)
	rpcOut, err = l.svcCtx.MsgRpc.GetPreviousMsgList(l.ctx, &message.GetPreviousMsgListIn{
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
