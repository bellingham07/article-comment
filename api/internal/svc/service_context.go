package svc

import (
	"article-comment/api/internal/config"
	"github.com/zeromicro/go-zero/core/stores/mon"
)

type ServiceContext struct {
	Config  config.Config
	Comment *mon.Model
}

func NewServiceContext(c config.Config) *ServiceContext {
	mongoUrl := InitMongo(c)
	return &ServiceContext{
		Config: c,
		// TODO 这里可以是把db也放入到配置文件中管理，上面的mongoUrl换成一个结构体，这里为了方便就不做了
		Comment: mon.MustNewModel(mongoUrl, "articledb", "comment"),
	}
}

func InitMongo(c config.Config) string {
	return c.Mongo.Url
}
