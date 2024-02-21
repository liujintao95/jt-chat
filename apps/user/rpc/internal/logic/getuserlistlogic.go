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

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserListLogic) GetUserList(in *pb.GetUserListIn) (*pb.GetUserListOut, error) {
	var (
		userList []*model.User
		err      error
		out      *pb.GetUserListOut
	)
	out = &pb.GetUserListOut{}
	userList, err = l.svcCtx.UserModel.FindPageByUidLike(l.ctx, in.Uid, in.Page, in.Size)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "模糊查询(%s)用户列表", in.Uid))
	}
	if len(userList) > 0 {
		for _, user := range userList {
			out.UserList = append(out.UserList, &pb.User{
				Uid:    user.Uid,
				Mobile: user.Mobile,
				Name:   user.Name,
				Avatar: user.Avatar.String,
			})
		}
	}
	out.Total, err = l.svcCtx.UserModel.FindCountByUidLike(l.ctx, in.Uid)
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "模糊查询(%s)用户列表总数", in.Uid))
	}
	return out, nil
}
