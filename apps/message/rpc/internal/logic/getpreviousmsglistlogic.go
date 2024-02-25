package logic

import (
	"context"

	"jt-chat/apps/message/rpc/internal/svc"
	"jt-chat/apps/message/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPreviousMsgListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPreviousMsgListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPreviousMsgListLogic {
	return &GetPreviousMsgListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPreviousMsgListLogic) GetPreviousMsgList(in *pb.GetPreviousMsgListIn) (*pb.GetPreviousMsgListOut, error) {
	// todo: add your logic here and delete this line

	return &pb.GetPreviousMsgListOut{}, nil
}
