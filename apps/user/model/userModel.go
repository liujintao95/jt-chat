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

		FindPageByUidLike(ctx context.Context, uidLike string, page, size int64) ([]*User, error)
		FindCountByUidLike(ctx context.Context, uidLike string) (int64, error)
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

func (m *customUserModel) FindPageByUidLike(ctx context.Context, uidLike string, page, size int64) ([]*User, error) {
	var (
		sqlFmt string
		resp   []*User
		offset int64
	)
	sqlFmt = `
	select %s from %s
	where uid like '%%?%%'
	and delete_at is null
	limit ?,?
	`
	offset = (page - 1) * size
	query := fmt.Sprintf(sqlFmt, userRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, uidLike, offset, size)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *customUserModel) FindCountByUidLike(ctx context.Context, uidLike string) (int64, error) {
	var (
		sqlFmt string
		resp   int64
	)
	sqlFmt = `
	select count(id) from %s
	where uid like '%%?%%'
	and delete_at is null
	`
	query := fmt.Sprintf(sqlFmt, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, uidLike)
	if err != nil {
		return 0, err
	}
	return resp, nil
}
