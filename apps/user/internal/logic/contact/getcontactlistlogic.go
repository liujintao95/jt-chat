package contact

import (
	"context"
	"github.com/pkg/errors"
	"jt-chat/apps/user/model"
	"jt-chat/common/constant"
	"jt-chat/common/ctxdata"
	"jt-chat/common/xerr"

	"jt-chat/apps/user/internal/svc"
	"jt-chat/apps/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContactListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetContactListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContactListLogic {
	return &GetContactListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetContactListLogic) GetContactList(req *types.GetContactListReq) (resp *types.GetContactListResp, err error) {
	var (
		uid         string
		contactList []*model.ContactOutput
	)
	uid = ctxdata.GetUidFromCtx(l.ctx)
	if req.NameOrObjectId == "" {
		contactList, err = l.svcCtx.UserContactModel.FindPageByUid(l.ctx, uid, req.Page, req.Size)
		resp.Total, err = l.svcCtx.UserContactModel.FindCountByUid(l.ctx, uid)
	} else {
		contactList, err = l.svcCtx.UserContactModel.FindPageByUidFuzzyInfo(l.ctx, uid, req.NameOrObjectId, req.Page, req.Size)
		resp.Total, err = l.svcCtx.UserContactModel.FindCountByUidFuzzyInfo(l.ctx, uid, req.NameOrObjectId)
	}
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, xerr.CustomErr(xerr.DbError, l.ctx, errors.Wrapf(err, "获取用户%s联系人列表", uid))
	}
	if len(contactList) > 0 {
		for _, contact := range contactList {
			var avatar string
			if contact.ContactType == constant.UserContactType {
				avatar = contact.UserAvatar.String
			} else {
				avatar = contact.GroupAvatar.String
			}
			resp.ContactList = append(resp.ContactList, types.Contact{
				ContactId:      contact.ContactId,
				ContactType:    contact.ContactType,
				Avatar:         avatar,
				NoteName:       contact.NoteName,
				LastMsgId:      contact.LastMsgId.Int64,
				LastMsgContent: contact.LastMsgContent.String,
				LastMsgTime:    contact.LastMsgTime.Time.Format("2006-01-02 15:04:05"),
			})
		}
	}
	return
}
