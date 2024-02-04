package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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
	}

	customUserContactModel struct {
		*defaultUserContactModel
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
