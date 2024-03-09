package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MessageModel = (*customMessageModel)(nil)

type (
	// MessageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageModel.
	MessageModel interface {
		messageModel
		withSession(session sqlx.Session) MessageModel
		FindOneByMsgId(ctx context.Context, msgId string) (*Message, error)
		FindNextPageByMsgId(ctx context.Context, msgId, uid, targetId string, size int64) ([]*Message, error)
		FindPreviousPageByMsgId(ctx context.Context, msgId, uid, targetId string, size int64) ([]*Message, error)
	}

	customMessageModel struct {
		*defaultMessageModel
	}
)

// NewMessageModel returns a model for the database table.
func NewMessageModel(conn sqlx.SqlConn) MessageModel {
	return &customMessageModel{
		defaultMessageModel: newMessageModel(conn),
	}
}

func (m *customMessageModel) withSession(session sqlx.Session) MessageModel {
	return NewMessageModel(sqlx.NewSqlConnFromSession(session))
}

func (m *defaultMessageModel) FindOneByMsgId(ctx context.Context, msgId string) (*Message, error) {
	query := fmt.Sprintf("select %s from %s where `msg_id` = ? limit 1", messageRows, m.table)
	var resp *Message
	err := m.conn.QueryRowCtx(ctx, resp, query, msgId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *customMessageModel) FindNextPageByMsgId(ctx context.Context, msgId, uid, targetId string, size int64) ([]*Message, error) {
	var (
		resp []*Message
	)
	query := fmt.Sprintf(`
	select %s from %s
	where msg_id > ?
	and (
		(
			from = ?
			and to = ?
		) or (
			from = ?
			and to = ?
		)
	)
	and delete_at is null
	order by id
	limit 1,?
	`, messageRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, msgId, uid, targetId, targetId, uid, size)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *customMessageModel) FindPreviousPageByMsgId(ctx context.Context, msgId, uid, targetId string, size int64) ([]*Message, error) {
	var (
		resp []*Message
	)
	query := fmt.Sprintf(`
	select %s from %s
	where msg_id < ?
	and (
		(
			from = ?
			and to = ?
		) or (
			from = ?
			and to = ?
		)
	)
	and delete_at is null
	order by id desc
	limit 1,?
	`, messageRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, msgId, uid, targetId, targetId, uid, size)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
