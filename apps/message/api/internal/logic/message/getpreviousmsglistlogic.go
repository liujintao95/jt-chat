package message

import (
	"context"

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
	// todo: add your service here and delete this line

	return
}
