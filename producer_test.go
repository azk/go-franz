package franz

import (
	"testing"
)

func TestNewProducerEmptyConfig(t *testing.T) {

	c := NewConfig()

	p, err := NewProducer(c)
	if p != nil || err != InvalidConfigurationError {
		t.Fail()
	}
}

