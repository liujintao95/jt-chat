package logic

import (
	"context"

	"jt-chat/apps/message/api/internal/svc"
	"jt-chat/apps/message/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageLogic {
	return &MessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageLogic) Message(req *types.Request) (resp *types.Response, err error) {
	// todo: add your service here and delete this line

	return
}
