package logic

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"jt-chat/apps/message/model"
	"jt-chat/common/xerr"

	"jt-chat/apps/message/rpc/internal/svc"
	"jt-chat/apps/message/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMsgLogic {
	return &CreateMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateMsgLogic) CreateMsg(in *pb.CreateMsgIn) (*pb.CreateMsgOut, error) {
	var (
		newMessage *model.Message
		err        error
	)
	newMessage = &model.Message{
		MsgId:         in.Msg.MsgId,
		TransportType: in.Msg.TransportType,
		From:          in.Msg.From,
		To:            in.Msg.To,
		ToType:        in.Msg.ToType,
		Content:       in.Msg.Content,
		ContentType:   in.Msg.ContentType,
		FilePath:      sql.NullString{String: in.Msg.FilePath, Valid: in.Msg.FilePath != ""},
		FileSuffix:    sql.NullString{String: in.Msg.FileSuffix, Valid: in.Msg.FileSuffix != ""},
	}
	_, err = l.svcCtx.MessageModel.Insert(l.ctx, newMessage)
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "创建信息[%+v]", newMessage))
	}
	return &pb.CreateMsgOut{}, nil
}
