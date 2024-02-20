package group

import (
	"context"
	"github.com/jinzhu/copier"
	"jt-chat/apps/user/api/internal/svc"
	"jt-chat/apps/user/api/internal/types"
	"jt-chat/apps/user/rpc/user"

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
		rpcOut *user.DeleteUserGroupMapOut
	)
	rpcOut, err = l.svcCtx.UserRpc.DeleteUserGroupMap(l.ctx, &user.DeleteUserGroupMapIn{
		Gid: req.Gid,
		Uid: req.Uid,
	})
	if err != nil {
		return nil, err
	}
	_ = copier.Copy(&resp, rpcOut)
	return
}
