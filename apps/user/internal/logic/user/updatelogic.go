package user

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"jt-chat/apps/user/internal/svc"
	"jt-chat/apps/user/internal/types"
	"jt-chat/apps/user/model"
	"jt-chat/common/ctxdata"
	"jt-chat/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateReq) (resp *types.UpdateResp, err error) {
	var (
		uid         string
		currentUser *model.User
	)
	uid = ctxdata.GetUidFromCtx(l.ctx)
	currentUser, err = l.svcCtx.UserModel.FindOneByUid(l.ctx, uid)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "获取登录用户%s信息", uid))
	}
	if currentUser == nil {
		return nil, xerr.CustomErr(xerr.UserNotExists, l.ctx, errors.New(fmt.Sprintf("用户%s不存在", uid)))
	}
	currentUser.Name = req.Name
	currentUser.Avatar = sql.NullString{String: req.Avatar, Valid: req.Avatar != ""}
	err = l.svcCtx.UserModel.Update(l.ctx, currentUser)
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "更新用户数据[%+v]", currentUser))
	}
	return
}
