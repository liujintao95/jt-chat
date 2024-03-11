package logic

import (
	"context"
	"github.com/pkg/errors"
	"jt-chat/apps/message/model"
	"jt-chat/common/ctxdata"
	"jt-chat/common/xerr"
	"os"

	"jt-chat/apps/message/rpc/internal/svc"
	"jt-chat/apps/message/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDownloadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadFileLogic {
	return &DownloadFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DownloadFileLogic) DownloadFile(in *pb.DownloadFileIn) (*pb.DownloadFileOut, error) {
	var (
		uid     string
		message *model.Message
		file    *model.File
		out     *pb.DownloadFileOut
		err     error
	)
	out = &pb.DownloadFileOut{}
	uid = ctxdata.GetUidFromCtx(l.ctx)
	message, err = l.svcCtx.MessageModel.FindOneByMsgId(l.ctx, in.MsgId)
	if message.From != uid && message.To != uid {
		return nil, xerr.CustomErr(xerr.PermissionError, l.ctx, errors.Wrapf(err, "用户%s没有权限查看信息%s", uid, in.MsgId))
	}
	if !message.FilePath.Valid {
		return nil, xerr.CustomErr(xerr.FileNotExists, l.ctx, errors.Wrapf(err, "信息%s中没有文件", in.MsgId))
	}
	file, err = l.svcCtx.FileModel.FindOneByMsgId(l.ctx, in.MsgId)
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "数据库获取消息%s的文件信息", in.MsgId))
	}
	out.Ext = file.Ext
	out.Name = file.Name
	out.Size = file.Size
	out.Data, err = os.ReadFile(message.FilePath.String)
	if err != nil {
		return nil, xerr.CustomErr(xerr.FileOpenErr, l.ctx, errors.Wrapf(err, "打开文件%s", message.FilePath.String))
	}
	return out, nil
}
