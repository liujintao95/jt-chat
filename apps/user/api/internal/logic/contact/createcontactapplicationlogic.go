package contact

import (
	"context"
	"github.com/jinzhu/copier"
	"jt-chat/apps/user/api/internal/svc"
	"jt-chat/apps/user/api/internal/types"
	"jt-chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateContactApplicationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateContactApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateContactApplicationLogic {
	return &CreateContactApplicationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateContactApplicationLogic) CreateContactApplication(req *types.CreateContactApplicationReq) (resp *types.CreateContactApplicationResp, err error) {
	var (
		rpcOut *user.CreateContactApplicationOut
	)
	rpcOut, err = l.svcCtx.UserRpc.CreateContactApplication(l.ctx, &user.CreateContactApplicationIn{
		ObjectId:    req.ObjectId,
		ContactType: req.ContactType,
		Notice:      req.Notice,
	})
	if err != nil {
		return nil, err
	}
	_ = copier.Copy(&resp, rpcOut)
	return
}
