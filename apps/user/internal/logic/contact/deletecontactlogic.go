package contact

import (
	"context"
	"github.com/pkg/errors"
	"jt-chat/apps/user/model"
	"jt-chat/common/constant"
	"jt-chat/common/ctxdata"
	"jt-chat/common/xerr"

	"jt-chat/apps/user/internal/svc"
	"jt-chat/apps/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteContactLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteContactLogic {
	return &DeleteContactLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteContactLogic) DeleteContact(req *types.DeleteContactReq) (resp *types.DeleteContactResp, err error) {
	// todo: add your logic here and delete this line
	var (
		uid         string
		userContact *model.UserContact
	)
	uid = ctxdata.GetUidFromCtx(l.ctx)
	userContact, err = l.svcCtx.UserContactModel.FindOneByUidObjectId(l.ctx, uid, req.ObjectId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "数据库查询用户%s的联系人%s失败", uid, req.ObjectId))
	}
	if userContact == nil {
		return nil, xerr.CustomErr(xerr.ContactNotExists, l.ctx, errors.Wrapf(err, "用户%s未找到联系人%s", uid, req.ObjectId))
	}

	err = l.svcCtx.UserContactModel.DeleteByUidObjectId(l.ctx, uid, req.ObjectId)
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "用户%s删除联系对象%s失败", uid, req.ObjectId))
	}
	if userContact.ContactType == constant.UserContactType {
		// 如果联系人对象是用户，则需要删除双方的联系人信息
		err = l.svcCtx.UserContactModel.DeleteByUidObjectId(l.ctx, req.ObjectId, uid)
		if err != nil {
			return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "用户%s删除联系对象%s失败", req.ObjectId, uid))
		}
	}
	return
}
