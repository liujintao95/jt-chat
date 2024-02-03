package auth

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

	"jt-chat/apps/user/internal/svc"
	"jt-chat/apps/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	var (
		user          *model.User
		registerLogic *RegisterLogic
		now           int64
	)
	user, err = l.svcCtx.UserModel.FindOneByMobile(l.ctx, req.Mobile)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "根据手机号%s查询用户信息", req.Mobile))
	}
	if user == nil {
		return nil, xerr.CustomErr(xerr.UserMobilePwdError, l.ctx, errors.Wrapf(err, "手机号%s未找到用户", req.Mobile))
	}
	if !(tool.Md5ByString(req.Password) == user.Password) {
		return nil, xerr.CustomErr(xerr.UserMobilePwdError, l.ctx, errors.New(fmt.Sprintf("手机号%s的密码%s错误", req.Mobile, req.Password)))
	}
	registerLogic = NewRegisterLogic(l.ctx, l.svcCtx)
	resp.AccessToken, err = registerLogic.GenerateToken(user.Uid)
	if err != nil {
		return nil, xerr.CustomErr(xerr.TokenGenerateError, l.ctx, errors.Wrapf(err, "用户%s生成toekn", user.Uid))
	}
	now = time.Now().Unix()
	resp.AccessExpire = now + l.svcCtx.Config.JwtAuth.AccessExpire
	resp.RefreshAfter = now + l.svcCtx.Config.JwtAuth.AccessExpire/2
	return
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
