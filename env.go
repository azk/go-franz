package franz

import (
	"os"
)

const (
	brokersEnv = "KAFKA_BROKERS"
)

type envProvider struct {}

func newEnv() *envProvider {
	return &envProvider{}
}

func (p *envProvider) PopulateConfig(conf *Config) (*Config, error) {

	conf.Brokers = os.Getenv(brokersEnv)

	return conf, nil
}



