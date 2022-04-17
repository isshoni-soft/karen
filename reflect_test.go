package karen

import (
	"reflect"
	"testing"
)

type testStruct struct {
	someField string `sometag:"thevalue"`
}

func TestFindFieldWithMatchingTag(t *testing.T) {
	test := testStruct{
		someField: "data",
	}
	field := FindFieldWithMatchingTag(reflect.ValueOf(test), "sometag", "thevalue")

	if field.String() != "data" {
		t.Fatal("did not get expected data from found field")
	}
}
