package contact

import (
	"context"
	"github.com/pkg/errors"
	"jt-chat/apps/user/model"
	"jt-chat/common/xerr"

	"jt-chat/apps/user/internal/svc"
	"jt-chat/apps/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupContactApplicationListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGroupContactApplicationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupContactApplicationListLogic {
	return &GetGroupContactApplicationListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupContactApplicationListLogic) GetGroupContactApplicationList(req *types.GetGroupContactApplicationListReq) (resp *types.GetGroupContactApplicationListResp, err error) {
	var (
		applications []*model.ContactApplication
	)
	applications, err = l.svcCtx.ContactApplicationModel.FindPageByObjectId(l.ctx, req.Gid, req.Page, req.Size)
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "获取群组%s申请信息列表", req.Gid))
	}
	resp.Total, err = l.svcCtx.ContactApplicationModel.FindCountByObjectId(l.ctx, req.Gid)
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "获取群组%s申请信息列表数量", req.Gid))
	}
	if len(applications) > 0 {
		for _, application := range applications {
			resp.ApplicationList = append(resp.ApplicationList, types.Application{
				ApplicationId: application.ApplicationId,
				Name:          application.Name,
				Avatar:        application.Avatar.String,
				ContactType:   application.ObjectType,
				Status:        application.Status,
				Notice:        application.Notes.String,
			})
		}
	}
	return
}
