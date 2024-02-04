package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"jt-chat/common/sessionctx"
)

var _ GroupModel = (*customGroupModel)(nil)

type (
	// GroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupModel.
	GroupModel interface {
		groupModel
		withSession(session sqlx.Session) GroupModel
		Insert(ctx context.Context, data *Group) (sql.Result, error)
	}

	customGroupModel struct {
		*defaultGroupModel
	}
)

// NewGroupModel returns a model for the database table.
func NewGroupModel(conn sqlx.SqlConn) GroupModel {
	return &customGroupModel{
		defaultGroupModel: newGroupModel(conn),
	}
}

func (m *customGroupModel) withSession(session sqlx.Session) GroupModel {
	return NewGroupModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customGroupModel) Insert(ctx context.Context, data *Group) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, groupRowsExpectAutoSet)
	ret, err := sessionctx.GetSession(ctx, m.conn).ExecCtx(ctx, query, data.Id, data.Gid, data.Name, data.Avatar, data.Notice, data.AdminUid, data.DeletedAt)
	return ret, err
}
