package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"jt-chat/apps/user/internal/config"
	"jt-chat/apps/user/model"
)

type ServiceContext struct {
	Config  config.Config
	SqlConn sqlx.SqlConn

	UserModel        model.UserModel
	GroupModel       model.GroupModel
	UserContactModel model.UserContactModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:  c,
		SqlConn: sqlConn,

		UserModel:        model.NewUserModel(sqlConn),
		GroupModel:       model.NewGroupModel(sqlConn),
		UserContactModel: model.NewUserContactModel(sqlConn),
	}
}
