package group

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"jt-chat/apps/user/model"
	"jt-chat/common/ctxdata"
	"jt-chat/common/xerr"

	"jt-chat/apps/user/internal/svc"
	"jt-chat/apps/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserGroupMapLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserGroupMapLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserGroupMapLogic {
	return &DeleteUserGroupMapLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserGroupMapLogic) DeleteUserGroupMap(req *types.DeleteUserGroupMapReq) (resp *types.DeleteUserGroupMapResp, err error) {
	var (
		uid   string
		group *model.Group
	)
	uid = ctxdata.GetUidFromCtx(l.ctx)
	if uid != req.Uid {
		// 管理员踢人
		group, err = l.svcCtx.GroupModel.FindOneByGid(l.ctx, req.Gid)
		if err != nil && !errors.Is(err, model.ErrNotFound) {
			return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "查询群聊组%s失败", req.Gid))
		}
		if group == nil {
			return nil, xerr.CustomErr(xerr.GroupNotExists, l.ctx, errors.Wrapf(err, "组ID%s未找到群聊组", req.Gid))
		}
		if group.AdminUid != uid {
			return nil, xerr.CustomErr(xerr.PermissionError, l.ctx, errors.New(fmt.Sprintf("用户%s在群聊组%s权限不够无法踢人", uid, req.Gid)))
		}
	}
	err = l.svcCtx.UserContactModel.DeleteByUidObjectId(l.ctx, req.Uid, req.Gid)
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "删除联系对象信息[%+v]失败", req))
	}
	return
}
