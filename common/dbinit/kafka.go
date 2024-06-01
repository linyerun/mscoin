package dbinit

import "github.com/segmentio/kafka-go"

func CreateKafkaTopicWriter(topic string, addresses []string) (writer *kafka.Writer) {
	writer = &kafka.Writer{
		Addr:         kafka.TCP(addresses...),
		RequiredAcks: kafka.RequireAll,
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{}, //将消息路由到接收到的数据量最少的分区
	}
	return
}
