package franz

import (
	"testing"
	"os"
)

func TestNewConfigFromEnvNoEnv(t *testing.T) {

	c, _ := NewConfigFromEnv()
	if c.Brokers != "" {
		t.Fail()
	}
}

func TestNewConfigFromEnv(t *testing.T) {
	val := "test1"

	os.Setenv(brokersEnv, val)

	c, _ := NewConfigFromEnv()
	if c.Brokers != val {
		t.Fail()
	}
}