package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"jt-chat/apps/user/model"
	"jt-chat/common/sessionctx"
	"jt-chat/common/xerr"

	"jt-chat/apps/user/rpc/internal/svc"
	"jt-chat/apps/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGroupLogic {
	return &DeleteGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteGroupLogic) DeleteGroup(in *pb.DeleteGroupIn) (*pb.DeleteGroupOut, error) {
	var (
		uid        string
		group      *model.Group
		sessionCtx context.Context
		err        error
	)
	group, err = l.svcCtx.GroupModel.FindOneByGid(l.ctx, in.Gid)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "查询群聊组%s失败", in.Gid))
	}
	if group == nil {
		return nil, xerr.CustomErr(xerr.GroupNotExists, l.ctx, errors.Wrapf(err, "组ID%s未找到群聊组", in.Gid))
	}
	if group.AdminUid != uid {
		return nil, xerr.CustomErr(xerr.PermissionError, l.ctx, errors.New(fmt.Sprintf("用户%s在群聊组%s无解散权限", uid, in.Gid)))
	}

	err = l.svcCtx.SqlConn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		sessionCtx = sessionctx.NewCtx(ctx, session)
		err = l.svcCtx.UserContactModel.DeleteByObjectId(sessionCtx, in.Gid)
		if err != nil {
			return errors.Wrapf(err, "删除群组%s所有联系人失败", in.Gid)
		}
		err = l.svcCtx.GroupModel.DeleteByGid(sessionCtx, in.Gid)
		if err != nil {
			return errors.Wrapf(err, "删除群组%s失败", in.Gid)
		}
		return nil
	})
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrap(err, "执行SQL事务失败"))
	}
	return &pb.DeleteGroupOut{}, nil
}
