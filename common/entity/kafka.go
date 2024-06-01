package entity

type KafkaData struct {
	Topic string
	Key   []byte
	Data  []byte
}
