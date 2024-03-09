package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ FileModel = (*customFileModel)(nil)

type (
	// FileModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFileModel.
	FileModel interface {
		fileModel
		withSession(session sqlx.Session) FileModel
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
