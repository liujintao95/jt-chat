package logic

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"jt-chat/apps/user/model"
	"jt-chat/common/ctxdata"
	"jt-chat/common/tool"
	"jt-chat/common/xerr"
	"time"

	"jt-chat/apps/user/rpc/internal/svc"
	"jt-chat/apps/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginIn) (*pb.LoginOut, error) {
	var (
		user          *model.User
		registerLogic *RegisterLogic
		now           int64
		err           error
		out           *pb.LoginOut
	)
	out = &pb.LoginOut{}
	user, err = l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "根据手机号%s查询用户信息", in.Mobile))
	}
	if user == nil {
		return nil, xerr.CustomErr(xerr.UserMobilePwdError, l.ctx, errors.Wrapf(err, "手机号%s未找到用户", in.Mobile))
	}
	if !(tool.Md5ByString(in.Password) == user.Password) {
		return nil, xerr.CustomErr(xerr.UserMobilePwdError, l.ctx, errors.New(fmt.Sprintf("手机号%s的密码%s错误", in.Mobile, in.Password)))
	}
	registerLogic = NewRegisterLogic(l.ctx, l.svcCtx)
	out.AccessToken, err = registerLogic.GenerateToken(user.Uid)
	if err != nil {
		return nil, xerr.CustomErr(xerr.TokenGenerateError, l.ctx, errors.Wrapf(err, "用户%s生成toekn", user.Uid))
	}
	now = time.Now().Unix()
	out.AccessExpire = now + l.svcCtx.Config.JwtAuth.AccessExpire
	out.RefreshAfter = now + l.svcCtx.Config.JwtAuth.AccessExpire/2
	return out, nil
}

func (l *RegisterLogic) GenerateToken(userId string) (string, error) {
	var (
		now          int64
		accessExpire int64
		accessToken  string
		err          error
	)
	now = time.Now().Unix()
	accessExpire = l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err = l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, userId, now, accessExpire)
	if err != nil {
		return "", xerr.CustomErr(xerr.TokenGenerateError, l.ctx, errors.New("GenerateToken err"))
	}

	return accessToken, nil
}

func (l *RegisterLogic) getJwtToken(secretKey, userId string, iat, seconds int64) (string, error) {
	var (
		claims jwt.MapClaims
		token  *jwt.Token
	)
	claims = make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[ctxdata.CtxKeyJwtUserId] = userId
	token = jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
