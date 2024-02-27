package logic

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"jt-chat/apps/user/model"
	"jt-chat/common/constant"
	"jt-chat/common/ctxdata"
	"jt-chat/common/sessionctx"
	"jt-chat/common/xerr"

	"jt-chat/apps/user/rpc/internal/svc"
	"jt-chat/apps/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateGroupLogic) CreateGroup(in *pb.CreateGroupIn) (*pb.CreateGroupOut, error) {
	var (
		uid            string
		newGroup       *model.Group
		newUserContact *model.UserContact
		sessionCtx     context.Context
		err            error
	)
	uid = ctxdata.GetUidFromCtx(l.ctx)
	newGroup = &model.Group{
		Gid:      uuid.New().String(),
		Name:     in.Name,
		Avatar:   in.Avatar,
		Notice:   sql.NullString{String: in.Avatar, Valid: in.Avatar != ""},
		AdminUid: uid,
	}
	newUserContact = &model.UserContact{
		ContactId:   uuid.New().String(),
		Uid:         uid,
		ObjectId:    newGroup.Gid,
		ContactType: constant.GroupContactType,
		NoteName:    newGroup.Name,
	}
	err = l.svcCtx.SqlConn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		sessionCtx = sessionctx.NewCtx(ctx, session)
		_, err = l.svcCtx.GroupModel.Insert(sessionCtx, newGroup)
		if err != nil {
			return errors.Wrapf(err, "创建群聊信息[%+v]", newGroup)
		}
		_, err = l.svcCtx.UserContactModel.Insert(sessionCtx, newUserContact)
		if err != nil {
			return errors.Wrapf(err, "创建联系人信息[%+v]", newUserContact)
		}
		return nil
	})
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrap(err, "执行SQL事务失败"))
	}
	return &pb.CreateGroupOut{Gid: newGroup.Gid}, nil
}
