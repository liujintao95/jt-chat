package contact

import (
	"context"
	"github.com/jinzhu/copier"
	"jt-chat/apps/user/api/internal/svc"
	"jt-chat/apps/user/api/internal/types"
	"jt-chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContactListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetContactListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContactListLogic {
	return &GetContactListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetContactListLogic) GetContactList(req *types.GetContactListReq) (resp *types.GetContactListResp, err error) {
	var (
		rpcOut *user.GetContactListOut
	)
	rpcOut, err = l.svcCtx.UserRpc.GetContactList(l.ctx, &user.GetContactListIn{
		NameOrObjectId: req.NameOrObjectId,
		Page:           req.Page,
		Size:           req.Size,
	})
	if err != nil {
		return nil, err
	}
	_ = copier.Copy(&resp, rpcOut)
	return
}
