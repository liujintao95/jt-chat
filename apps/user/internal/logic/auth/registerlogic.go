package auth

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"jt-chat/apps/user/internal/svc"
	"jt-chat/apps/user/internal/types"
	"jt-chat/apps/user/model"
	"jt-chat/common/constant"
	"jt-chat/common/tool"
	"jt-chat/common/xerr"
	"regexp"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	var (
		mobileRegexp *regexp.Regexp
		user         *model.User
	)
	mobileRegexp = regexp.MustCompile(constant.MobileRegex)
	if !mobileRegexp.MatchString(req.Mobile) {
		return nil, xerr.CustomErr(xerr.RequestParamError, l.ctx, errors.New(fmt.Sprintf("手机号%s格式", req.Mobile)))
	}
	user, err = l.svcCtx.UserModel.FindOneByMobile(l.ctx, req.Mobile)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "查询手机号%s", req.Mobile))
	}
	if user != nil {
		return nil, xerr.CustomErr(xerr.UserAlreadyExists, l.ctx, errors.Wrapf(err, "手机号%s已被注册", req.Mobile))
	}
	user = new(model.User)
	user.Uid = uuid.New().String()
	user.Name = req.Name
	user.Mobile = req.Mobile
	user.Password = tool.Md5ByString(req.Password)
	_, err = l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "用户注册[%+v]", user))
	}
	resp.Mobile = req.Mobile
	return
}
