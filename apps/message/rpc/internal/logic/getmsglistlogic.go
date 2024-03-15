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

type GetMsgListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMsgListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMsgListLogic {
	return &GetMsgListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMsgListLogic) GetMsgList(in *pb.GetMsgListIn) (*pb.GetMsgListOut, error) {
	// todo: add your logic here and delete this line
	var (
		uid      string
		messages []*model.Message
		out      *pb.GetMsgListOut
		err      error
	)
	out = &pb.GetMsgListOut{}
	uid = ctxdata.GetUidFromCtx(l.ctx)
	messages, err = l.svcCtx.MessageModel.FindPageByContentLike(l.ctx, in.ContentLike, uid, in.TargetId, in.Page, in.Size)
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "模糊查询%s消息列表", in.ContentLike))
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
	out.Total, err = l.svcCtx.MessageModel.FindCountByContentLike(l.ctx, in.ContentLike, uid, in.TargetId, in.Page, in.Size)
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "模糊查询%s消息列表总数", in.ContentLike))
	}
	return out, nil
}
