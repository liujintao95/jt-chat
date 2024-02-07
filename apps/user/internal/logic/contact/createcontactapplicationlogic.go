package contact

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"jt-chat/apps/user/model"
	"jt-chat/common/constant"
	"jt-chat/common/ctxdata"
	"jt-chat/common/xerr"

	"jt-chat/apps/user/internal/svc"
	"jt-chat/apps/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateContactApplicationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateContactApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateContactApplicationLogic {
	return &CreateContactApplicationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateContactApplicationLogic) CreateContactApplication(req *types.CreateContactApplicationReq) (resp *types.CreateContactApplicationResp, err error) {
	// todo: add your logic here and delete this line
	var (
		uid            string
		userContact    *model.UserContact
		currentUser    *model.User
		newApplication *model.ContactApplication
	)
	uid = ctxdata.GetUidFromCtx(l.ctx)
	userContact, err = l.svcCtx.UserContactModel.FindOneByUidObjectId(l.ctx, uid, req.ObjectId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "数据库查询用户%s的联系人%s失败", uid, req.ObjectId))
	}
	if userContact != nil {
		return nil, xerr.CustomErr(xerr.ContactAlreadyExists, l.ctx, errors.Wrapf(err, "用户%s已填加联系人%s，不能重复添加", uid, req.ObjectId))
	}
	currentUser, err = l.svcCtx.UserModel.FindOneByUid(l.ctx, uid)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "数据库查询用户%s失败", uid))
	}

	newApplication = &model.ContactApplication{
		ApplicationId: uuid.New().String(),
		Uid:           uid,
		Name:          currentUser.Name,
		Avatar:        currentUser.Avatar,
		ObjectId:      req.ObjectId,
		ObjectType:    req.ContactType,
		Notes:         sql.NullString{String: req.Notice, Valid: req.Notice != ""},
		Status:        constant.WaitApplicationStatus,
	}
	_, err = l.svcCtx.ContactApplicationModel.Insert(l.ctx, newApplication)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "添加联系人申请[%+v]失败", newApplication))
	}
	return
}
