package logic

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"jt-chat/apps/user/model"
	"jt-chat/common/constant"
	"jt-chat/common/tool"
	"jt-chat/common/xerr"
	"regexp"

	"jt-chat/apps/user/rpc/internal/svc"
	"jt-chat/apps/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterIn) (*pb.RegisterOut, error) {
	var (
		mobileRegexp *regexp.Regexp
		user         *model.User
		err          error
	)
	mobileRegexp = regexp.MustCompile(constant.MobileRegex)
	if !mobileRegexp.MatchString(in.Mobile) {
		return nil, xerr.CustomErr(xerr.RequestParamError, l.ctx, errors.New(fmt.Sprintf("手机号%s格式", in.Mobile)))
	}
	user, err = l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "查询手机号%s", in.Mobile))
	}
	if user != nil {
		return nil, xerr.CustomErr(xerr.UserAlreadyExists, l.ctx, errors.Wrapf(err, "手机号%s已被注册", in.Mobile))
	}
	user = new(model.User)
	user.Uid = uuid.New().String()
	user.Name = in.Name
	user.Mobile = in.Mobile
	user.Password = tool.Md5ByString(in.Password)
	_, err = l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "用户注册[%+v]", user))
	}
	return &pb.RegisterOut{Mobile: in.Mobile}, nil
}
