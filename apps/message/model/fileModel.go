package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FileModel = (*customFileModel)(nil)

type (
	// FileModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFileModel.
	FileModel interface {
		fileModel
		withSession(session sqlx.Session) FileModel
		FindOneByMsgId(ctx context.Context, msgId string) (*File, error)
	}

	customFileModel struct {
		*defaultFileModel
	}
)

// NewFileModel returns a model1 for the database table.
func NewFileModel(conn sqlx.SqlConn) FileModel {
	return &customFileModel{
		defaultFileModel: newFileModel(conn),
	}
}

func (m *customFileModel) withSession(session sqlx.Session) FileModel {
	return NewFileModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customFileModel) FindOneByMsgId(ctx context.Context, msgId string) (*File, error) {
	query := fmt.Sprintf("select %s from %s where `msg_id` = ? limit 1", fileRows, m.table)
	var resp *File
	err := m.conn.QueryRowCtx(ctx, resp, query, msgId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
