package contact

import (
	"context"

	"jt-chat/apps/user/internal/svc"
	"jt-chat/apps/user/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
