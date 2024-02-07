package contact

import (
	"context"
	"github.com/pkg/errors"
	"jt-chat/apps/user/model"
	"jt-chat/common/ctxdata"
	"jt-chat/common/xerr"

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
	var (
		uid          string
		applications []*model.ContactApplication
	)
	uid = ctxdata.GetUidFromCtx(l.ctx)
	applications, err = l.svcCtx.ContactApplicationModel.FindPageByObjectId(l.ctx, uid, req.Page, req.Size)
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "获取联系人%s申请信息列表", uid))
	}
	resp.Total, err = l.svcCtx.ContactApplicationModel.FindCountByObjectId(l.ctx, uid)
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "获取联系人%s申请信息列表数量", uid))
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
