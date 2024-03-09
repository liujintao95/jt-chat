package logic

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
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
		newMsg *model.Message
		err    error
	)
	newMsg = &model.Message{
		MsgId:         uuid.New().String(),
		TransportType: in.TransportType,
		From:          in.From,
		To:            in.To,
		ToType:        in.ToType,
		Content:       in.Content,
		ContentType:   in.ContentType,
		FilePath:      sql.NullString{String: in.FilePath, Valid: in.FilePath != ""},
		FileExt:       sql.NullString{String: in.FileExt, Valid: in.FileExt != ""},
	}
	_, err = l.svcCtx.MessageModel.Insert(l.ctx, newMsg)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "保存消息信息[%+v]", newMsg))
	}
	return &pb.CreateMsgOut{}, nil
}
