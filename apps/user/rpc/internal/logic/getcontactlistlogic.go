package logic

import (
	"context"
	"github.com/pkg/errors"
	"jt-chat/apps/user/model"
	"jt-chat/common/constant"
	"jt-chat/common/ctxdata"
	"jt-chat/common/xerr"

	"jt-chat/apps/user/rpc/internal/svc"
	"jt-chat/apps/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContactListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetContactListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContactListLogic {
	return &GetContactListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetContactListLogic) GetContactList(in *pb.GetContactListIn) (*pb.GetContactListOut, error) {
	var (
		uid         string
		contactList []*model.ContactOutput
		err         error
		out         *pb.GetContactListOut
	)
	out = &pb.GetContactListOut{}
	uid = ctxdata.GetUidFromCtx(l.ctx)
	if in.NameOrObjectId == "" {
		contactList, err = l.svcCtx.UserContactModel.FindPageByUid(l.ctx, uid, in.Page, in.Size)
		out.Total, err = l.svcCtx.UserContactModel.FindCountByUid(l.ctx, uid)
	} else {
		contactList, err = l.svcCtx.UserContactModel.FindPageByUidFuzzyInfo(l.ctx, uid, in.NameOrObjectId, in.Page, in.Size)
		out.Total, err = l.svcCtx.UserContactModel.FindCountByUidFuzzyInfo(l.ctx, uid, in.NameOrObjectId)
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
			out.ContactList = append(out.ContactList, &pb.Contact{
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
	return out, nil
}
