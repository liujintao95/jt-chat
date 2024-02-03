package contact

import (
	"context"

	"jt-chat/apps/user/internal/svc"
	"jt-chat/apps/user/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
