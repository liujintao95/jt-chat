package user

import (
	"context"
	"github.com/pkg/errors"
	"jt-chat/apps/user/model"
	"jt-chat/common/xerr"

	"jt-chat/apps/user/internal/svc"
	"jt-chat/apps/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListLogic {
	return &GetListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetListLogic) GetList(req *types.GetListReq) (resp *types.GetListResp, err error) {
	var (
		userList []*model.User
	)
	userList, err = l.svcCtx.UserModel.FindPageByNameOrUid(l.ctx, req.NameOrUid, req.Page, req.Size)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "模糊查询(%s)用户列表", req.NameOrUid))
	}
	if len(userList) > 0 {
		for _, user := range userList {
			resp.UserList = append(resp.UserList, types.User{
				Uid:    user.Uid,
				Mobile: user.Mobile,
				Name:   user.Name,
				Avatar: user.Avatar.String,
			})
		}
	}
	resp.Total, err = l.svcCtx.UserModel.FindCountByNameOrUid(l.ctx, req.NameOrUid)
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "模糊查询(%s)用户列表总数", req.NameOrUid))
	}
	return
}
