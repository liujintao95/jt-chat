// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	contactApplicationFieldNames          = builder.RawFieldNames(&ContactApplication{})
	contactApplicationRows                = strings.Join(contactApplicationFieldNames, ",")
	contactApplicationRowsExpectAutoSet   = strings.Join(stringx.Remove(contactApplicationFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	contactApplicationRowsWithPlaceHolder = strings.Join(stringx.Remove(contactApplicationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	contactApplicationModel interface {
		Insert(ctx context.Context, data *ContactApplication) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ContactApplication, error)
		FindOneByApplicationId(ctx context.Context, applicationId string) (*ContactApplication, error)
		Update(ctx context.Context, data *ContactApplication) error
		Delete(ctx context.Context, id int64) error
	}

	defaultContactApplicationModel struct {
		conn  sqlx.SqlConn
		table string
	}

	ContactApplication struct {
		Id            int64          `db:"id"`
		ApplicationId string         `db:"application_id"` // application_id
		Uid           string         `db:"uid"`            // uid
		Name          string         `db:"name"`           // 用户名
		Avatar        sql.NullString `db:"avatar"`         // 头像
		ObjectId      string         `db:"object_id"`      // 请求对象的gid或uid
		ObjectType    int64          `db:"object_type"`    // 请求对象的类型
		Notes         sql.NullString `db:"notes"`          // 备注
		Status        int64          `db:"status"`         // 状态
		DeletedAt     sql.NullTime   `db:"deleted_at"`     // 删除时间
		CreatedAt     time.Time      `db:"created_at"`     // 创建时间
		UpdatedAt     time.Time      `db:"updated_at"`     // 更新时间
	}
)

func newContactApplicationModel(conn sqlx.SqlConn) *defaultContactApplicationModel {
	return &defaultContactApplicationModel{
		conn:  conn,
		table: "`contact_application`",
	}
}

func (m *defaultContactApplicationModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultContactApplicationModel) FindOne(ctx context.Context, id int64) (*ContactApplication, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", contactApplicationRows, m.table)
	var resp ContactApplication
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultContactApplicationModel) FindOneByApplicationId(ctx context.Context, applicationId string) (*ContactApplication, error) {
	var resp ContactApplication
	query := fmt.Sprintf("select %s from %s where `application_id` = ? limit 1", contactApplicationRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, applicationId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultContactApplicationModel) Insert(ctx context.Context, data *ContactApplication) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, contactApplicationRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.ApplicationId, data.Uid, data.Name, data.Avatar, data.ObjectId, data.ObjectType, data.Notes, data.Status, data.DeletedAt)
	return ret, err
}

func (m *defaultContactApplicationModel) Update(ctx context.Context, newData *ContactApplication) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, contactApplicationRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.ApplicationId, newData.Uid, newData.Name, newData.Avatar, newData.ObjectId, newData.ObjectType, newData.Notes, newData.Status, newData.DeletedAt, newData.Id)
	return err
}

func (m *defaultContactApplicationModel) tableName() string {
	return m.table
}
