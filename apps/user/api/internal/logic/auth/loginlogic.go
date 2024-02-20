package auth

import (
	"context"
	"github.com/jinzhu/copier"
	"jt-chat/apps/user/api/internal/svc"
	"jt-chat/apps/user/api/internal/types"
	"jt-chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	var (
		rpcOut *user.LoginOut
	)
	rpcOut, err = l.svcCtx.UserRpc.Login(l.ctx, &user.LoginIn{
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	_ = copier.Copy(&resp, rpcOut)
	return
}
