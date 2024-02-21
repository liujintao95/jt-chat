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

type GetGroupUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupUserListLogic {
	return &GetGroupUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupUserListLogic) GetGroupUserList(in *pb.GetGroupUserListIn) (*pb.GetGroupUserListOut, error) {
	var (
		out           *pb.GetGroupUserListOut
		groupUserList []*model.GroupUserOutput
		err           error
	)
	out = &pb.GetGroupUserListOut{}
	if in.NameOrUid != "" {
		groupUserList, err = l.svcCtx.UserContactModel.FindUserPageByGidFuzzyInfo(l.ctx, in.Gid, in.NameOrUid, in.Page, in.Size)
		out.Total, err = l.svcCtx.UserContactModel.FindUserCountByGidFuzzyInfo(l.ctx, in.Gid, in.NameOrUid)
	} else {
		groupUserList, err = l.svcCtx.UserContactModel.FindUserPageByGid(l.ctx, in.Gid, in.Page, in.Size)
		out.Total, err = l.svcCtx.UserContactModel.FindUserCountByGid(l.ctx, in.Gid)
	}
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "获取群%s用户列表", in.Gid))
	}
	if len(groupUserList) > 0 {
		for _, user := range groupUserList {
			out.GroupUserList = append(out.GroupUserList, &pb.GroupUser{
				Uid:      user.Uid,
				Avatar:   user.Avatar.String,
				NoteName: user.NoteName,
				NickName: user.NickName,
			})
		}
	}
	return out, nil
}
