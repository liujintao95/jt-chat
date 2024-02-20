package logic

import (
	"context"
	"github.com/pkg/errors"
	"jt-chat/apps/user/model"
	"jt-chat/common/xerr"

	"jt-chat/apps/user/rpc/internal/svc"
	"jt-chat/apps/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupContactApplicationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupContactApplicationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupContactApplicationListLogic {
	return &GetGroupContactApplicationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupContactApplicationListLogic) GetGroupContactApplicationList(in *pb.GetGroupContactApplicationListIn) (*pb.GetGroupContactApplicationListOut, error) {
	var (
		applications []*model.ContactApplication
		err          error
		out          *pb.GetGroupContactApplicationListOut
	)
	out = &pb.GetGroupContactApplicationListOut{}
	applications, err = l.svcCtx.ContactApplicationModel.FindPageByObjectId(l.ctx, in.Gid, in.Page, in.Size)
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "获取群组%s申请信息列表", in.Gid))
	}
	out.Total, err = l.svcCtx.ContactApplicationModel.FindCountByObjectId(l.ctx, in.Gid)
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "获取群组%s申请信息列表数量", in.Gid))
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
