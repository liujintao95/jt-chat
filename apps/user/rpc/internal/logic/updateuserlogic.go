package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"jt-chat/apps/user/model"
	"jt-chat/common/ctxdata"
	"jt-chat/common/xerr"

	"jt-chat/apps/user/rpc/internal/svc"
	"jt-chat/apps/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *pb.UpdateIn) (*pb.UpdateOut, error) {
	var (
		uid         string
		currentUser *model.User
		err         error
	)
	uid = ctxdata.GetUidFromCtx(l.ctx)
	currentUser, err = l.svcCtx.UserModel.FindOneByUid(l.ctx, uid)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "获取登录用户%s信息", uid))
	}
	if currentUser == nil {
		return nil, xerr.CustomErr(xerr.UserNotExists, l.ctx, errors.New(fmt.Sprintf("用户%s不存在", uid)))
	}
	currentUser.Name = in.Name
	currentUser.Avatar = sql.NullString{String: in.Avatar, Valid: in.Avatar != ""}
	err = l.svcCtx.UserModel.Update(l.ctx, currentUser)
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "更新用户数据[%+v]", currentUser))
	}
	return &pb.UpdateOut{}, nil
}
