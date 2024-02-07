package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/timex"
	"jt-chat/common/constant"
	"jt-chat/common/sessionctx"
)

var _ UserContactModel = (*customUserContactModel)(nil)

type (
	// UserContactModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserContactModel.
	UserContactModel interface {
		userContactModel
		withSession(session sqlx.Session) UserContactModel
		Insert(ctx context.Context, data *UserContact) (sql.Result, error)
		DeleteByUidObjectId(ctx context.Context, uid, objectId string) error
		DeleteByObjectId(ctx context.Context, objectId string) error
		FindPageByUid(ctx context.Context, uid string, page, size int64) ([]*ContactOutput, error)
		FindCountByUid(ctx context.Context, uid string) (int64, error)
		FindOneByUidObjectId(ctx context.Context, uid, objectId string) (*UserContact, error)
	}

	customUserContactModel struct {
		*defaultUserContactModel
	}

	ContactOutput struct {
		ContactId      string         `db:"contact_id"`       // contact_id
		ContactType    int64          `db:"contact_type"`     // 联系人类型
		NoteName       string         `db:"note_name"`        // 备注名称
		LastMsgId      sql.NullInt64  `db:"last_msg_id"`      // 最后一条消息ID
		LastMsgContent sql.NullString `db:"last_msg_content"` // 最后一条消息内容
		LastMsgTime    sql.NullTime   `db:"last_msg_time"`    // 最后一条消息时间
		UserAvatar     sql.NullString `db:"user_avatar"`      // 用户头像
		GroupAvatar    sql.NullString `db:"group_avatar"`     // 群组头像
	}
)

// NewUserContactModel returns a model for the database table.
func NewUserContactModel(conn sqlx.SqlConn) UserContactModel {
	return &customUserContactModel{
		defaultUserContactModel: newUserContactModel(conn),
	}
}

func (m *customUserContactModel) withSession(session sqlx.Session) UserContactModel {
	return NewUserContactModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customUserContactModel) Insert(ctx context.Context, data *UserContact) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userContactRowsExpectAutoSet)
	ret, err := sessionctx.GetSession(ctx, m.conn).ExecCtx(ctx, query, data.Id, data.ContactId, data.Uid, data.ObjectId, data.ContactType, data.NoteName, data.LastMsgId, data.LastMsgContent, data.LastMsgTime, data.DeletedAt)
	return ret, err
}

func (m *customUserContactModel) DeleteByUidObjectId(ctx context.Context, uid, objectId string) error {
	query := fmt.Sprintf("update %s set `delete_at` = ? where `uid` = ? and `object_id` = ?", m.table)
	_, err := sessionctx.GetSession(ctx, m.conn).ExecCtx(ctx, query, timex.Now(), uid, objectId)
	return err
}

func (m *customUserContactModel) DeleteByObjectId(ctx context.Context, objectId string) error {
	query := fmt.Sprintf("update %s set `delete_at` = ? where `object_id` = ?", m.table)
	_, err := sessionctx.GetSession(ctx, m.conn).ExecCtx(ctx, query, timex.Now(), objectId)
	return err
}

func (m *customUserContactModel) FindPageByUid(ctx context.Context, uid string, page, size int64) ([]*ContactOutput, error) {
	var (
		offset int64
		sqlFmt string
		resp   []*ContactOutput
	)
	sqlFmt = `
	select contact_id, contact_type, note_name, last_msg_id, last_msg_content, 
	       last_msg_time, u.avatar as user_avatar, g.avatar as group_avatar 
	from user_contact as uc
	left join user as u
	on uc.object_id = u.uid
	and uc.ContactType = %s
	left join group as g
	on uc.object_id = g.gid
	and uc.ContactType = %s
	where uid = ?
	and delete_at is null
	limit ?,?
	`
	offset = (page - 1) * size
	query := fmt.Sprintf(sqlFmt, userRows, constant.UserContactType, constant.GroupContactType)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, uid, offset, size)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *customUserContactModel) FindCountByUid(ctx context.Context, uid string) (int64, error) {
	var (
		sqlFmt string
		resp   int64
	)
	sqlFmt = `
	select count(uc.id)
	from user_contact as uc
	left join user as u
	on uc.object_id = u.uid
	and uc.ContactType = %s
	left join group as g
	on uc.object_id = g.gid
	and uc.ContactType = %s
	where uid = ?
	and delete_at is null
	`
	query := fmt.Sprintf(sqlFmt, userRows, constant.UserContactType, constant.GroupContactType)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, uid)
	if err != nil {
		return 0, err
	}
	return resp, nil
}

func (m *customUserContactModel) FindOneByUidObjectId(ctx context.Context, uid, objectId string) (*UserContact, error) {
	var (
		resp *UserContact
	)
	query := fmt.Sprintf("select %s from %s where `uid` = ? and `object_id` = ? and delete_at is null limit 1", userContactRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, uid, objectId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
