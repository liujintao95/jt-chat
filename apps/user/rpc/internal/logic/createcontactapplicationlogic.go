package logic

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"jt-chat/apps/user/model"
	"jt-chat/common/constant"
	"jt-chat/common/ctxdata"
	"jt-chat/common/xerr"

	"jt-chat/apps/user/rpc/internal/svc"
	"jt-chat/apps/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateContactApplicationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateContactApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateContactApplicationLogic {
	return &CreateContactApplicationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateContactApplicationLogic) CreateContactApplication(in *pb.CreateContactApplicationIn) (*pb.CreateContactApplicationOut, error) {
	var (
		uid            string
		userContact    *model.UserContact
		currentUser    *model.User
		newApplication *model.ContactApplication
		err            error
	)
	uid = ctxdata.GetUidFromCtx(l.ctx)
	userContact, err = l.svcCtx.UserContactModel.FindOneByUidObjectId(l.ctx, uid, in.ObjectId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "数据库查询用户%s的联系人%s", uid, in.ObjectId))
	}
	if userContact != nil {
		return nil, xerr.CustomErr(xerr.ContactAlreadyExists, l.ctx, errors.Wrapf(err, "用户%s已填加联系人%s，不能重复添加", uid, in.ObjectId))
	}
	currentUser, err = l.svcCtx.UserModel.FindOneByUid(l.ctx, uid)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "数据库查询用户%s", uid))
	}

	newApplication = &model.ContactApplication{
		ApplicationId: uuid.New().String(),
		Uid:           uid,
		Name:          currentUser.Name,
		Avatar:        currentUser.Avatar,
		ObjectId:      in.ObjectId,
		ObjectType:    in.ContactType,
		Notes:         sql.NullString{String: in.Notice, Valid: in.Notice != ""},
		Status:        constant.WaitApplicationStatus,
	}
	_, err = l.svcCtx.ContactApplicationModel.Insert(l.ctx, newApplication)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "添加联系人申请[%+v]失败", newApplication))
	}
	return &pb.CreateContactApplicationOut{ApplicationId: newApplication.ApplicationId}, nil
}
