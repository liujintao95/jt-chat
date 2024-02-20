package group

import (
	"context"
	"github.com/jinzhu/copier"
	"jt-chat/apps/user/api/internal/svc"
	"jt-chat/apps/user/api/internal/types"
	"jt-chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateGroupLogic) CreateGroup(req *types.CreateGroupReq) (resp *types.CreateGroupResp, err error) {
	var (
		rpcOut *user.CreateGroupOut
	)
	rpcOut, err = l.svcCtx.UserRpc.CreateGroup(l.ctx, &user.CreateGroupIn{
		Name:   req.Name,
		Avatar: req.Avatar,
		Notice: req.Notice,
	})
	if err != nil {
		return nil, err
	}
	_ = copier.Copy(&resp, rpcOut)
	return
}
