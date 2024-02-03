package contact

import (
	"context"

	"jt-chat/apps/user/internal/svc"
	"jt-chat/apps/user/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
