package logic

import (
	"context"
	"github.com/pkg/errors"
	"jt-chat/apps/user/model"
	"jt-chat/common/constant"
	"jt-chat/common/ctxdata"
	"jt-chat/common/xerr"

	"jt-chat/apps/user/rpc/internal/svc"
	"jt-chat/apps/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteContactLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteContactLogic {
	return &DeleteContactLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteContactLogic) DeleteContact(in *pb.DeleteContactIn) (*pb.DeleteContactOut, error) {
	var (
		uid         string
		userContact *model.UserContact
		err         error
	)
	uid = ctxdata.GetUidFromCtx(l.ctx)
	userContact, err = l.svcCtx.UserContactModel.FindOneByUidObjectId(l.ctx, uid, in.ObjectId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "数据库查询用户%s的联系人%s", uid, in.ObjectId))
	}
	if userContact == nil {
		return nil, xerr.CustomErr(xerr.ContactNotExists, l.ctx, errors.Wrapf(err, "用户%s未找到联系人%s", uid, in.ObjectId))
	}

	err = l.svcCtx.UserContactModel.DeleteByUidObjectId(l.ctx, uid, in.ObjectId)
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "用户%s删除联系对象%s", uid, in.ObjectId))
	}
	if userContact.ContactType == constant.UserContactType {
		// 如果联系人对象是用户，则需要删除双方的联系人信息
		err = l.svcCtx.UserContactModel.DeleteByUidObjectId(l.ctx, in.ObjectId, uid)
		if err != nil {
			return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "用户%s删除联系对象%s", in.ObjectId, uid))
		}
	}
	return &pb.DeleteContactOut{}, nil
}
