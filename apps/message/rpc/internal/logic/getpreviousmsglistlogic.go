package logic

import (
	"context"
	"github.com/pkg/errors"
	"jt-chat/apps/message/model"
	"jt-chat/common/ctxdata"
	"jt-chat/common/xerr"

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
	var (
		uid        string
		currentMsg *model.Message
		messages   []*model.Message
		out        *pb.GetPreviousMsgListOut
		err        error
	)
	out = &pb.GetPreviousMsgListOut{}
	uid = ctxdata.GetUidFromCtx(l.ctx)
	currentMsg, err = l.svcCtx.MessageModel.FindOneByMsgId(l.ctx, in.MsgId)
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "数据库查询消息%s", in.MsgId))
	}
	if currentMsg == nil {
		return nil, xerr.CustomErr(xerr.MsgNotExists, l.ctx, errors.Wrapf(err, "未找到消息%s", in.MsgId))
	}

	messages, err = l.svcCtx.MessageModel.FindPreviousPageByMsgId(l.ctx, in.MsgId, uid, in.TargetId, in.Size)
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "数据库获取消息%s的上一页", in.MsgId))
	}
	if len(messages) > 0 {
		for _, msg := range messages {
			out.MessageList = append(out.MessageList, &pb.Message{
				MsgId:         msg.MsgId,
				TransportType: msg.TransportType,
				From:          msg.From,
				To:            msg.To,
				ToType:        msg.ToType,
				Content:       msg.Content,
				ContentType:   msg.ContentType,
				FileExt:       msg.FileExt.String,
				FilePath:      msg.FilePath.String,
				CreatedAt:     msg.CreatedAt.Format("2006-01-02 15:04:05"),
			})
		}
	}
	return out, nil
}
