package franz

import "testing"

func TestNegotiate(t *testing.T) {

	p, _ := Negotiate()

	switch p.(type) {
	case *envProvider:
		return
	default:
		t.Fail()
	}

}