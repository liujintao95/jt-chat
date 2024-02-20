package contact

import (
	"context"
	"github.com/jinzhu/copier"
	"jt-chat/apps/user/api/internal/svc"
	"jt-chat/apps/user/api/internal/types"
	"jt-chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteContactLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteContactLogic {
	return &DeleteContactLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteContactLogic) DeleteContact(req *types.DeleteContactReq) (resp *types.DeleteContactResp, err error) {
	var (
		rpcOut *user.CreateContactApplicationOut
	)
	rpcOut, err = l.svcCtx.UserRpc.CreateContactApplication(l.ctx, &user.CreateContactApplicationIn{
		ObjectId: req.ObjectId,
	})
	if err != nil {
		return nil, err
	}
	_ = copier.Copy(&resp, rpcOut)
	return
}
