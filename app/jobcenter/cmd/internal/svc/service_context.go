package svc

import (
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/mongo"
	"mscoin/app/jobcenter/cmd/internal/config"
	"mscoin/common/dbinit"
)

type ServiceContext struct {
	Config            config.Config
	MongoClient       *mongo.Client
	KafkaWriterCreate func(topic string) *kafka.Writer
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化MongoDB远程连接
	mongoDBCli, err := dbinit.CreateMongoDBClient(c.Mongo.Username, c.Mongo.Password, c.Mongo.Url)
	if err != nil {
		logx.Error(err)
	}

	return &ServiceContext{
		Config:      c,
		MongoClient: mongoDBCli,
		KafkaWriterCreate: func(topic string) *kafka.Writer {
			return dbinit.CreateKafkaTopicWriter(topic, c.Kafka.Addresses)
		},
	}
}
