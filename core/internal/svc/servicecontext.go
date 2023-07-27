package svc

import (
	"core/internal/middleware"
	"github.com/zeromicro/go-zero/rest"

	"core/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	AccessLog rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	//conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:    c,
		AccessLog: middleware.NewAccessLogMiddleware().Handle,
	}
}
