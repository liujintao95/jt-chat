package group

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"jt-chat/apps/user/model"
	"jt-chat/common/xerr"

	"jt-chat/apps/user/internal/svc"
	"jt-chat/apps/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupLogic {
	return &UpdateGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateGroupLogic) UpdateGroup(req *types.UpdateGroupReq) (resp *types.UpdateGroupResp, err error) {
	var (
		uid   string
		group *model.Group
	)
	group, err = l.svcCtx.GroupModel.FindOneByGid(l.ctx, req.Gid)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "查询群聊组%s失败", req.Gid))
	}
	if group == nil {
		return nil, xerr.CustomErr(xerr.GroupNotExists, l.ctx, errors.Wrapf(err, "组ID%s未找到群聊组", req.Gid))
	}
	if group.AdminUid != uid {
		return nil, xerr.CustomErr(xerr.PermissionError, l.ctx, errors.New(fmt.Sprintf("用户%s在群聊组%s无修改权限", uid, req.Gid)))
	}

	group.Name = req.Name
	group.Avatar = req.Avatar
	group.Notice = sql.NullString{String: req.Notice, Valid: req.Notice != ""}
	group.AdminUid = req.AdminUid
	err = l.svcCtx.GroupModel.Update(l.ctx, group)
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "修改联系对象信息[%+v]失败", group))
	}
	return
}
