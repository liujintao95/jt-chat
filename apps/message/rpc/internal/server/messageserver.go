// Code generated by goctl. DO NOT EDIT.
// Source: message.proto

package server

import (
	"context"

	"jt-chat/apps/message/rpc/internal/logic"
	"jt-chat/apps/message/rpc/internal/svc"
	"jt-chat/apps/message/rpc/pb"
)

type MessageServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedMessageServer
}

func NewMessageServer(svcCtx *svc.ServiceContext) *MessageServer {
	return &MessageServer{
		svcCtx: svcCtx,
	}
}

func (s *MessageServer) CreateMsg(ctx context.Context, in *pb.CreateMsgIn) (*pb.CreateMsgOut, error) {
	l := logic.NewCreateMsgLogic(ctx, s.svcCtx)
	return l.CreateMsg(in)
}

func (s *MessageServer) GetNextMsgList(ctx context.Context, in *pb.GetNextMsgListIn) (*pb.GetNextMsgListOut, error) {
	l := logic.NewGetNextMsgListLogic(ctx, s.svcCtx)
	return l.GetNextMsgList(in)
}

func (s *MessageServer) GetPreviousMsgList(ctx context.Context, in *pb.GetPreviousMsgListIn) (*pb.GetPreviousMsgListOut, error) {
	l := logic.NewGetPreviousMsgListLogic(ctx, s.svcCtx)
	return l.GetPreviousMsgList(in)
}

func (s *MessageServer) GetMsgList(ctx context.Context, in *pb.GetMsgListIn) (*pb.GetMsgListOut, error) {
	l := logic.NewGetMsgListLogic(ctx, s.svcCtx)
	return l.GetMsgList(in)
}

func (s *MessageServer) DownloadFile(ctx context.Context, in *pb.DownloadFileIn) (*pb.DownloadFileOut, error) {
	l := logic.NewDownloadFileLogic(ctx, s.svcCtx)
	return l.DownloadFile(in)
}

func (s *MessageServer) UploadFile(ctx context.Context, in *pb.UploadFileIn) (*pb.UploadFileOut, error) {
	l := logic.NewUploadFileLogic(ctx, s.svcCtx)
	return l.UploadFile(in)
}
