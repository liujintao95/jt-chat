package group

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"jt-chat/apps/user/model"
	"jt-chat/common/sessionctx"
	"jt-chat/common/xerr"

	"jt-chat/apps/user/internal/svc"
	"jt-chat/apps/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGroupLogic {
	return &DeleteGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteGroupLogic) DeleteGroup(req *types.DeleteGroupReq) (resp *types.DeleteGroupResp, err error) {
	var (
		uid        string
		group      *model.Group
		sessionCtx context.Context
	)
	group, err = l.svcCtx.GroupModel.FindOneByGid(l.ctx, req.Gid)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "查询群聊组%s失败", req.Gid))
	}
	if group == nil {
		return nil, xerr.CustomErr(xerr.GroupNotExists, l.ctx, errors.Wrapf(err, "组ID%s未找到群聊组", req.Gid))
	}
	if group.AdminUid != uid {
		return nil, xerr.CustomErr(xerr.PermissionError, l.ctx, errors.New(fmt.Sprintf("用户%s在群聊组%s无解散权限", uid, req.Gid)))
	}

	err = l.svcCtx.SqlConn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		sessionCtx = sessionctx.NewCtx(ctx, session)
		err = l.svcCtx.UserContactModel.DeleteByObjectId(sessionCtx, req.Gid)
		if err != nil {
			return errors.Wrapf(err, "删除群组%s所有联系人失败", req.Gid)
		}
		err = l.svcCtx.GroupModel.DeleteByGid(sessionCtx, req.Gid)
		if err != nil {
			return errors.Wrapf(err, "删除群组%s失败", req.Gid)
		}
		return nil
	})
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrap(err, "执行SQL事务失败"))
	}
	return
}
