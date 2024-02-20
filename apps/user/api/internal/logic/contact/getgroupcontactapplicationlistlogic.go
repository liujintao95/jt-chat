package contact

import (
	"context"
	"github.com/jinzhu/copier"
	"jt-chat/apps/user/api/internal/svc"
	"jt-chat/apps/user/api/internal/types"
	"jt-chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupContactApplicationListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGroupContactApplicationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupContactApplicationListLogic {
	return &GetGroupContactApplicationListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupContactApplicationListLogic) GetGroupContactApplicationList(req *types.GetGroupContactApplicationListReq) (resp *types.GetGroupContactApplicationListResp, err error) {
	var (
		rpcOut *user.GetGroupContactApplicationListOut
	)
	rpcOut, err = l.svcCtx.UserRpc.GetGroupContactApplicationList(l.ctx, &user.GetGroupContactApplicationListIn{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	_ = copier.Copy(&resp, rpcOut)
	return
}
