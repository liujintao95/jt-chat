package logic

import (
	"context"
	"github.com/pkg/errors"
	"jt-chat/apps/user/model"
	"jt-chat/common/ctxdata"
	"jt-chat/common/xerr"

	"jt-chat/apps/user/rpc/internal/svc"
	"jt-chat/apps/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContactApplicationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetContactApplicationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContactApplicationListLogic {
	return &GetContactApplicationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetContactApplicationListLogic) GetContactApplicationList(in *pb.GetContactApplicationListIn) (*pb.GetContactApplicationListOut, error) {
	var (
		uid          string
		applications []*model.ContactApplication
		err          error
		out          *pb.GetContactApplicationListOut
	)
	out = &pb.GetContactApplicationListOut{}
	uid = ctxdata.GetUidFromCtx(l.ctx)
	applications, err = l.svcCtx.ContactApplicationModel.FindPageByObjectId(l.ctx, uid, in.Page, in.Size)
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "获取联系人%s申请信息列表", uid))
	}
	out.Total, err = l.svcCtx.ContactApplicationModel.FindCountByObjectId(l.ctx, uid)
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "获取联系人%s申请信息列表数量", uid))
	}
	if len(applications) > 0 {
		for _, application := range applications {
			out.ApplicationList = append(out.ApplicationList, &pb.Application{
				ApplicationId: application.ApplicationId,
				Name:          application.Name,
				Avatar:        application.Avatar.String,
				ContactType:   application.ObjectType,
				Status:        application.Status,
				Notice:        application.Notes.String,
			})
		}
	}
	return out, nil
}
