package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"jt-chat/apps/user/model"
	"jt-chat/common/ctxdata"
	"jt-chat/common/xerr"

	"jt-chat/apps/user/rpc/internal/svc"
	"jt-chat/apps/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserGroupMapLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserGroupMapLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserGroupMapLogic {
	return &DeleteUserGroupMapLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserGroupMapLogic) DeleteUserGroupMap(in *pb.DeleteUserGroupMapIn) (*pb.DeleteUserGroupMapOut, error) {
	var (
		uid   string
		group *model.Group
		err   error
	)
	uid = ctxdata.GetUidFromCtx(l.ctx)
	if uid != in.Uid {
		// 管理员踢人
		group, err = l.svcCtx.GroupModel.FindOneByGid(l.ctx, in.Gid)
		if err != nil && !errors.Is(err, model.ErrNotFound) {
			return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "查询群聊组%s失败", in.Gid))
		}
		if group == nil {
			return nil, xerr.CustomErr(xerr.GroupNotExists, l.ctx, errors.Wrapf(err, "组ID%s未找到群聊组", in.Gid))
		}
		if group.AdminUid != uid {
			return nil, xerr.CustomErr(xerr.PermissionError, l.ctx, errors.New(fmt.Sprintf("用户%s在群聊组%s权限不够无法踢人", uid, in.Gid)))
		}
	}
	err = l.svcCtx.UserContactModel.DeleteByUidObjectId(l.ctx, in.Uid, in.Gid)
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "删除联系对象信息[%+v]失败", in))
	}

	return &pb.DeleteUserGroupMapOut{}, nil
}
