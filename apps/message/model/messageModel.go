package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ MessageModel = (*customMessageModel)(nil)

type (
	// MessageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageModel.
	MessageModel interface {
		messageModel
		withSession(session sqlx.Session) MessageModel
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
