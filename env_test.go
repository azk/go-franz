package franz

import (
	"testing"
	"os"
)

func TestEnvProvider_PopulateConfigNoEnv(t *testing.T) {

	os.Clearenv()

	e := newEnv()
	c := NewConfig()

	c, _ = e.PopulateConfig(c)

	if c.Brokers != "" {
		t.Fail()
	}
}

func TestEnvProvider_PopulateConfig(t *testing.T) {

	val := "test1"

	os.Setenv(brokersEnv, val)

	e := newEnv()
	c := NewConfig()

	c, _ = e.PopulateConfig(c)

	if c.Brokers != val {
		t.Fail()
	}
}
