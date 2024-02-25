package logic

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"jt-chat/apps/user/model"
	"jt-chat/common/ctxdata"
	"jt-chat/common/xerr"

	"jt-chat/apps/user/rpc/internal/svc"
	"jt-chat/apps/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateContactLastMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateContactLastMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateContactLastMsgLogic {
	return &UpdateContactLastMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateContactLastMsgLogic) UpdateContactLastMsg(in *pb.UpdateContactLastMsgIn) (*pb.UpdateContactLastMsgOut, error) {
	// todo: add your logic here and delete this line
	var (
		uid         string
		userContact *model.UserContact
		err         error
	)
	uid = ctxdata.GetUidFromCtx(l.ctx)
	userContact, err = l.svcCtx.UserContactModel.FindOneByUidObjectId(l.ctx, uid, in.ObjectId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "数据库查询用户%s联系人信息%s失败", uid, in.ObjectId))
	}
	if userContact == nil {
		return nil, xerr.CustomErr(xerr.ContactNotExists, l.ctx, errors.Wrapf(err, "用户%s未找到联系人%s", uid, in.ObjectId))
	}
	userContact.LastMsgId = sql.NullInt64{Int64: in.LastMsgId, Valid: in.LastMsgId != 0}
	userContact.LastMsgContent = sql.NullString{String: in.LastMsgContent, Valid: in.LastMsgContent != ""}
	userContact.LastMsgTime = sql.NullTime{Time: in.LastMsgTime.AsTime(), Valid: in.LastMsgTime.IsValid()}
	err = l.svcCtx.UserContactModel.Update(l.ctx, userContact)
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "修改联系人信息[%+v]", userContact))
	}
	return &pb.UpdateContactLastMsgOut{}, nil
}
