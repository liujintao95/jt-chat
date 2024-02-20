package contact

import (
	"context"
	"github.com/jinzhu/copier"
	"jt-chat/apps/user/api/internal/svc"
	"jt-chat/apps/user/api/internal/types"
	"jt-chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateContactApplicationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateContactApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateContactApplicationLogic {
	return &UpdateContactApplicationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateContactApplicationLogic) UpdateContactApplication(req *types.UpdateContactApplicationReq) (resp *types.UpdateContactApplicationResp, err error) {
	var (
		rpcOut *user.UpdateContactApplicationOut
	)
	rpcOut, err = l.svcCtx.UserRpc.UpdateContactApplication(l.ctx, &user.UpdateContactApplicationIn{
		ApplicationId: req.ApplicationId,
		Status:        req.Status,
	})
	if err != nil {
		return nil, err
	}
	_ = copier.Copy(&resp, rpcOut)
	return
}
