package config

type Config struct {
	KLineConf struct {
		ApiKey    string
		SecretKey string
		Pass      string
		Host      string
		Proxy     string
	}
	Mongo struct {
		Url      string
		Username string
		Password string
		Database string
	}
	Kafka struct {
		Addresses []string
	}
}
