package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ContactApplicationModel = (*customContactApplicationModel)(nil)

type (
	// ContactApplicationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customContactApplicationModel.
	ContactApplicationModel interface {
		contactApplicationModel
		withSession(session sqlx.Session) ContactApplicationModel
		FindPageByObjectId(ctx context.Context, objectId string, page, size int64) ([]*ContactApplication, error)
		FindCountByObjectId(ctx context.Context, objectId string) (int64, error)
	}

	customContactApplicationModel struct {
		*defaultContactApplicationModel
	}
)

// NewContactApplicationModel returns a model for the database table.
func NewContactApplicationModel(conn sqlx.SqlConn) ContactApplicationModel {
	return &customContactApplicationModel{
		defaultContactApplicationModel: newContactApplicationModel(conn),
	}
}

func (m *customContactApplicationModel) withSession(session sqlx.Session) ContactApplicationModel {
	return NewContactApplicationModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customContactApplicationModel) FindPageByObjectId(ctx context.Context, objectId string, page, size int64) ([]*ContactApplication, error) {
	var (
		resp   []*ContactApplication
		offset int64
	)
	offset = (page - 1) * size
	query := fmt.Sprintf("select %s from %s where object_id = ? and delete_at is nulllimit ?,?", userRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, objectId, offset, size)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *customContactApplicationModel) FindCountByObjectId(ctx context.Context, objectId string) (int64, error) {
	var (
		resp int64
	)
	query := fmt.Sprintf("select Count(id) from %s where object_id = ? and delete_at is nulllimit ?,?", m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, objectId)
	if err != nil {
		return 0, err
	}
	return resp, nil
}
