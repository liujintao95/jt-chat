package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"jt-chat/apps/user/model"
	"jt-chat/common/xerr"

	"jt-chat/apps/user/rpc/internal/svc"
	"jt-chat/apps/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupLogic {
	return &UpdateGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGroupLogic) UpdateGroup(in *pb.UpdateGroupIn) (*pb.UpdateGroupOut, error) {
	var (
		uid   string
		group *model.Group
		err   error
	)
	group, err = l.svcCtx.GroupModel.FindOneByGid(l.ctx, in.Gid)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "查询群聊组%s失败", in.Gid))
	}
	if group == nil {
		return nil, xerr.CustomErr(xerr.GroupNotExists, l.ctx, errors.Wrapf(err, "组ID%s未找到群聊组", in.Gid))
	}
	if group.AdminUid != uid {
		return nil, xerr.CustomErr(xerr.PermissionError, l.ctx, errors.New(fmt.Sprintf("用户%s在群聊组%s无修改权限", uid, in.Gid)))
	}

	group.Name = in.Name
	group.Avatar = in.Avatar
	group.Notice = sql.NullString{String: in.Notice, Valid: in.Notice != ""}
	group.AdminUid = in.AdminUid
	err = l.svcCtx.GroupModel.Update(l.ctx, group)
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "修改联系对象信息[%+v]失败", group))
	}
	return &pb.UpdateGroupOut{}, nil
}
