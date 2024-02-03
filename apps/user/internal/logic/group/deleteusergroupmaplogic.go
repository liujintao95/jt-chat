package group

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
