package contact

import (
	"context"

	"jt-chat/apps/user/internal/svc"
	"jt-chat/apps/user/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
