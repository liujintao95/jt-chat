package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ContactApplicationModel = (*customContactApplicationModel)(nil)

type (
	// ContactApplicationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customContactApplicationModel.
	ContactApplicationModel interface {
		contactApplicationModel
		withSession(session sqlx.Session) ContactApplicationModel
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
