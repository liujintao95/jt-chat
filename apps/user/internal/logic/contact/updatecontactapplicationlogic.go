package contact

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"jt-chat/apps/user/model"
	"jt-chat/common/constant"
	"jt-chat/common/ctxdata"
	"jt-chat/common/sessionctx"
	"jt-chat/common/xerr"

	"jt-chat/apps/user/internal/svc"
	"jt-chat/apps/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateContactApplicationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateContactApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateContactApplicationLogic {
	return &UpdateContactApplicationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateContactApplicationLogic) UpdateContactApplication(req *types.UpdateContactApplicationReq) (resp *types.UpdateContactApplicationResp, err error) {
	var (
		uid         string
		application *model.ContactApplication
		currentUser *model.User
		group       *model.Group
		sessionCtx  context.Context
		newContact  *model.UserContact
	)
	uid = ctxdata.GetUidFromCtx(l.ctx)
	application, err = l.svcCtx.ContactApplicationModel.FindOneByApplicationId(l.ctx, req.ApplicationId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "数据库查询联系人申请%s失败", req.ApplicationId))
	}
	if application == nil {
		return nil, xerr.CustomErr(xerr.ApplicationNotExists, l.ctx, errors.Wrapf(err, "未找到联系人申请%s", req.ApplicationId))
	}

	if application.ObjectType == constant.UserContactType {
		// 判断自己是否是联系人申请的对象
		currentUser, err = l.svcCtx.UserModel.FindOneByUid(l.ctx, uid)
		if err != nil && !errors.Is(err, model.ErrNotFound) {
			return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "数据库查询当前用户%s失败", uid))
		}
		if group == nil {
			return nil, xerr.CustomErr(xerr.GroupNotExists, l.ctx, errors.Wrapf(err, "未找到当前用户%s", uid))
		}
		if application.ObjectId != currentUser.Uid {
			return nil, xerr.CustomErr(xerr.PermissionError, l.ctx, errors.Wrapf(err, "用户%s不是联系人申请%s的请求对象", uid, req.ApplicationId))
		}
	} else if application.ObjectType == constant.GroupContactType {
		// 判断自己是否是联系人申请群组的管理员
		group, err = l.svcCtx.GroupModel.FindOneByGid(l.ctx, application.ObjectId)
		if err != nil && !errors.Is(err, model.ErrNotFound) {
			return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "数据库查询群组%s失败", application.ObjectId))
		}
		if group == nil {
			return nil, xerr.CustomErr(xerr.GroupNotExists, l.ctx, errors.Wrapf(err, "未找到群组%s", application.ObjectId))
		}
		if group.AdminUid != uid {
			return nil, xerr.CustomErr(xerr.PermissionError, l.ctx, errors.Wrapf(err, "用户%s群%s的管理员", uid, group.Gid))
		}
	}

	err = l.svcCtx.SqlConn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		sessionCtx = sessionctx.NewCtx(ctx, session)

		application.Status = req.Status
		err = l.svcCtx.ContactApplicationModel.Update(sessionCtx, application)
		if err != nil {
			return xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "修改联系人申请信息[%+v]", application))
		}
		if req.Status == constant.AgreeApplicationStatus {
			// 如果同意则需要添加联系人信息
			if application.ObjectType == constant.UserContactType {
				newContact = &model.UserContact{
					ContactId:   uuid.New().String(),
					Uid:         application.Uid,
					ObjectId:    currentUser.Uid,
					ContactType: application.ObjectType,
					NoteName:    currentUser.Name,
					NickName:    currentUser.Name,
				}
				_, err = l.svcCtx.UserContactModel.Insert(sessionCtx, newContact)
				if err != nil {
					return xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "创建联系人信息[%+v]", newContact))
				}
				newContact = &model.UserContact{
					ContactId:   uuid.New().String(),
					Uid:         currentUser.Uid,
					ObjectId:    application.Uid,
					ContactType: application.ObjectType,
					NoteName:    application.Name,
					NickName:    application.Name,
				}
				_, err = l.svcCtx.UserContactModel.Insert(sessionCtx, newContact)
				if err != nil {
					return xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "创建联系人信息[%+v]", newContact))
				}
			} else if application.ObjectType == constant.GroupContactType {
				newContact = &model.UserContact{
					ContactId:   uuid.New().String(),
					Uid:         application.Uid,
					ObjectId:    application.ObjectId,
					ContactType: application.ObjectType,
					NoteName:    group.Name,
					NickName:    group.Name,
				}
				_, err = l.svcCtx.UserContactModel.Insert(sessionCtx, newContact)
				if err != nil {
					return xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "创建联系人信息[%+v]", newContact))
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrap(err, "执行SQL事务失败"))
	}
	return
}
