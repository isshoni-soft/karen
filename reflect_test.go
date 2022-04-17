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

func TestIsEitherKind(t *testing.T) {
	testPtr := &testStruct{
		someField: "data",
	}

	test := testStruct{
		someField: "data",
	}

	if IsEitherKind(reflect.ValueOf(test), reflect.Pointer, reflect.Interface) {
		t.Fatal("detected ", test, " as either pointer or interface")
	}

	if !IsEitherKind(reflect.ValueOf(testPtr), reflect.Pointer, reflect.Interface) {
		t.Fatal(test, " did not detect as either pointer or interface")
	}
}
