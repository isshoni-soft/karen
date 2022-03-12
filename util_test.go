package karen

import "testing"

func TestZero(t *testing.T) {
	var str string

	if str != Zero[string]() {
		t.Fatal("zero does not match default string")
	}
}
