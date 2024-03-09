package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"jt-chat/apps/message/model"
	"jt-chat/apps/message/rpc/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	SqlConn sqlx.SqlConn

	MessageModel model.MessageModel
	FileModel    model.FileModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:  c,
		SqlConn: sqlConn,

		MessageModel: model.NewMessageModel(sqlConn),
		FileModel:    model.NewFileModel(sqlConn),
	}
}
