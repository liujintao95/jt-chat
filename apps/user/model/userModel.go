package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		withSession(session sqlx.Session) UserModel

		FindPageByNameOrUid(ctx context.Context, nameOrUid string, page, size int64) ([]*User, error)
		FindCountByNameOrUid(ctx context.Context, nameOrUid string) (int64, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn),
	}
}

func (m *customUserModel) withSession(session sqlx.Session) UserModel {
	return NewUserModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customUserModel) FindPageByNameOrUid(ctx context.Context, nameOrUid string, page, size int64) ([]*User, error) {
	var (
		sql    string
		resp   []*User
		offset int64
	)
	sql = `
	select %s from %s
	where (
	    name = '%%?%%' 
		or uid like '%%?%%'
	) 
	and delete_at is null
	limit %d,%d
	`
	offset = (page - 1) * size
	query := fmt.Sprintf(sql, userRows, m.table, offset, size)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, nameOrUid, nameOrUid)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *customUserModel) FindCountByNameOrUid(ctx context.Context, nameOrUid string) (int64, error) {
	var (
		sql  string
		resp int64
	)
	sql = `
	select count(id) from %s
	where (
	    name = '%%?%%' 
		or uid like '%%?%%'
	) 
	and delete_at is null
	`
	query := fmt.Sprintf(sql, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, nameOrUid, nameOrUid)
	if err != nil {
		return 0, err
	}
	return resp, nil
}
