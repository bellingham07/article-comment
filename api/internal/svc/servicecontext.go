package svc

import (
	"article-comment/api/internal/config"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceContext struct {
	Config config.Config
	Mongo  mongo.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Mongo:  mongo.NewClientEncryption(),
	}
}

func InitMongo(c config.Config) (mongo *mon.Model) {
	mongo = mon.MustNewModel("mongodb://root:caojinbo@42.194.238.75:14001", "db", "collection")
	return
}
