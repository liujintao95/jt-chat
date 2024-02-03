package contact

import (
	"context"

	"jt-chat/apps/user/internal/svc"
	"jt-chat/apps/user/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
