package logic

import (
	"context"

	"jt-chat/apps/message/rpc/internal/svc"
	"jt-chat/apps/message/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNextMsgListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNextMsgListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNextMsgListLogic {
	return &GetNextMsgListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetNextMsgListLogic) GetNextMsgList(in *pb.GetNextMsgListIn) (*pb.GetNextMsgListOut, error) {
	// todo: add your logic here and delete this line

	return &pb.GetNextMsgListOut{}, nil
}
