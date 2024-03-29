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
		FindPageByUidFuzzyInfo(ctx context.Context, uid, nameOrObjectId string, page, size int64) ([]*ContactOutput, error)
		FindCountByUidFuzzyInfo(ctx context.Context, uid, nameOrObjectId string) (int64, error)
		FindUserPageByGidFuzzyInfo(ctx context.Context, gid, fuzzyInfo string, page, size int64) ([]*GroupUserOutput, error)
		FindUserCountByGidFuzzyInfo(ctx context.Context, gid, fuzzyInfo string) (int64, error)
		FindUserPageByGid(ctx context.Context, gid string, page, size int64) ([]*GroupUserOutput, error)
		FindUserCountByGid(ctx context.Context, gid string) (int64, error)
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

	GroupUserOutput struct {
		Uid      string         `db:"uid"`       // uid
		NoteName string         `db:"note_name"` // 备注名称
		NickName string         `db:"nick_name"` // 备注名称
		Avatar   sql.NullString `db:"avatar"`    // 头像
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
		resp   []*ContactOutput
	)
	offset = (page - 1) * size
	query := fmt.Sprintf(`
	select contact_id, contact_type, note_name, last_msg_id, last_msg_content, 
	       last_msg_time, u.avatar as user_avatar, g.avatar as group_avatar 
	from user_contact as uc
	left join user as u
	on uc.object_id = u.uid
	and uc.ContactType = %d
	and u.delete_at is null
	left join group as g
	on uc.object_id = g.gid
	and uc.ContactType = %d
	and g.delete_at is null
	where uc.uid = ?
	and uc.delete_at is null
	limit ?,?
	`, constant.UserContactType, constant.GroupContactType)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, uid, offset, size)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *customUserContactModel) FindCountByUid(ctx context.Context, uid string) (int64, error) {
	var (
		resp int64
	)
	query := fmt.Sprintf(`
	select count(uc.id)
	from user_contact as uc
	left join user as u
	on uc.object_id = u.uid
	and uc.ContactType = %d
	and u.delete_at is null
	left join group as g
	on uc.object_id = g.gid
	and uc.ContactType = %d
	and g.delete_at is null
	where uc.uid = ?
	and uc.delete_at is null
	`, constant.UserContactType, constant.GroupContactType)
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
	err := m.conn.QueryRowsCtx(ctx, resp, query, uid, objectId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *customUserContactModel) FindPageByUidFuzzyInfo(ctx context.Context, uid, fuzzyInfo string, page, size int64) ([]*ContactOutput, error) {
	var (
		offset int64
		resp   []*ContactOutput
	)
	offset = (page - 1) * size
	query := fmt.Sprintf(`
	select contact_id, contact_type, note_name, last_msg_id, last_msg_content, 
	       last_msg_time, u.avatar as user_avatar, g.avatar as group_avatar 
	from user_contact as uc
	left join user as u
	on uc.object_id = u.uid
	and uc.ContactType = %d
	and u.delete_at is null
	left join group as g
	on uc.object_id = g.gid
	and uc.ContactType = %d
	and g.delete_at is null
	where uc.uid = ?
	and (
		object_id like '%%?%%'
		or note_name like '%%?%%'
		or nick_name like '%%?%%'
	)
	and uc.delete_at is null
	limit ?,?
	`, constant.UserContactType, constant.GroupContactType)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, uid, fuzzyInfo, fuzzyInfo, fuzzyInfo, offset, size)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *customUserContactModel) FindCountByUidFuzzyInfo(ctx context.Context, uid, fuzzyInfo string) (int64, error) {
	var (
		resp int64
	)
	query := fmt.Sprintf(`
	select count(uc.id)
	from user_contact as uc
	left join user as u
	on uc.object_id = u.uid
	and uc.ContactType = %d
	and u.delete_at is null
	left join group as g
	on uc.object_id = g.gid
	and uc.ContactType = %d
	and g.delete_at is null
	where uc.uid = ?
	and (
		object_id like '%%?%%'
		or note_name like '%%?%%'
		or nick_name like '%%?%%'
	)
	and uc.delete_at is null
	limit ?,?
	`, constant.UserContactType, constant.GroupContactType)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, uid, fuzzyInfo, fuzzyInfo, fuzzyInfo)
	if err != nil {
		return 0, err
	}
	return resp, nil
}

func (m *customUserContactModel) FindUserPageByGidFuzzyInfo(ctx context.Context, gid, fuzzyInfo string, page, size int64) ([]*GroupUserOutput, error) {
	var (
		offset int64
		resp   []*GroupUserOutput
	)
	offset = (page - 1) * size
	query := fmt.Sprintf(`
	select u.uid, u.avatar, note_name, nick_name
	from user_contact as uc
	inner join user as u
	on uc.uid = u.uid
	and uc.ContactType = %d
	where object_id = ?
	and (
		uc.uid like '%%?%%'
		or note_name like '%%?%%'
		or nick_name like '%%?%%'
	)
	and uc.delete_at is null
	and u.delete_at is null
	limit ?,?
	`, constant.GroupContactType)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, gid, fuzzyInfo, fuzzyInfo, fuzzyInfo, offset, size)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *customUserContactModel) FindUserCountByGidFuzzyInfo(ctx context.Context, gid, fuzzyInfo string) (int64, error) {
	var (
		resp int64
	)
	query := fmt.Sprintf(`
	select count(u.id)
	from user_contact as uc
	inner join user as u
	on uc.uid = u.uid
	and uc.ContactType = %d
	where object_id = ?
	and (
		uc.uid like '%%?%%'
		or note_name like '%%?%%'
		or nick_name like '%%?%%'
	)
	and uc.delete_at is null
	and u.delete_at is null
	`, constant.GroupContactType)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, gid, fuzzyInfo, fuzzyInfo, fuzzyInfo)
	if err != nil {
		return 0, err
	}
	return resp, nil
}

func (m *customUserContactModel) FindUserPageByGid(ctx context.Context, gid string, page, size int64) ([]*GroupUserOutput, error) {
	var (
		offset int64
		resp   []*GroupUserOutput
	)
	offset = (page - 1) * size
	query := fmt.Sprintf(`
	select u.uid, u.avatar, note_name, nick_name
	from user_contact as uc
	inner join user as u
	on uc.uid = u.uid
	and uc.ContactType = %d
	where object_id = ?
	and uc.delete_at is null
	and u.delete_at is null
	limit ?,?
	`, constant.GroupContactType)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, gid, offset, size)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *customUserContactModel) FindUserCountByGid(ctx context.Context, gid string) (int64, error) {
	var (
		resp int64
	)
	query := fmt.Sprintf(`
	select count(u.id)
	from user_contact as uc
	inner join user as u
	on uc.uid = u.uid
	and uc.ContactType = %d
	where object_id = ?
	and uc.delete_at is null
	and u.delete_at is null
	`, constant.GroupContactType)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, gid)
	if err != nil {
		return 0, err
	}
	return resp, nil
}
