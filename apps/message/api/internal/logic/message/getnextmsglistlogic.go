package message

import (
	"context"

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
	// todo: add your service here and delete this line

	return
}
