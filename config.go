package franz

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ConfigMap kafka.ConfigMap

type Config struct {
	Brokers string

	Properties   ConfigMap
}


func NewConfig() *Config {
	return &Config{}
}

func NewConfigFromEnv() (*Config, error) {

	provider, err := Negotiate()
	if err != nil {
		return nil, err
	}

	return provider.PopulateConfig(NewConfig())
}
