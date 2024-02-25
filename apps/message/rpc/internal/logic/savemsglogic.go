package logic

import (
	"context"

	"jt-chat/apps/message/rpc/internal/svc"
	"jt-chat/apps/message/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveMsgLogic {
	return &SaveMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveMsgLogic) SaveMsg(in *pb.SaveMsgIn) (*pb.SaveMsgOut, error) {
	// todo: add your logic here and delete this line

	return &pb.SaveMsgOut{}, nil
}
