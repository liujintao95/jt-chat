package contact

import (
	"context"
	"github.com/jinzhu/copier"
	"jt-chat/apps/user/api/internal/svc"
	"jt-chat/apps/user/api/internal/types"
	"jt-chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContactApplicationListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetContactApplicationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContactApplicationListLogic {
	return &GetContactApplicationListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetContactApplicationListLogic) GetContactApplicationList(req *types.GetContactApplicationListReq) (resp *types.GetContactApplicationListResp, err error) {
	var (
		rpcOut *user.GetContactApplicationListOut
	)
	rpcOut, err = l.svcCtx.UserRpc.GetContactApplicationList(l.ctx, &user.GetContactApplicationListIn{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	_ = copier.Copy(&resp, rpcOut)
	return
}
