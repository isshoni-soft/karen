package karen

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type testInterface interface {
	Test()
}

type testStruct struct {
	someField string `sometag:"thevalue"`
}

func (t testStruct) Test() {}

func TestExecuteOnField(t *testing.T) {
	assert := assert.New(t)

	test := testStruct{
		someField: "value",
	}

	ExecuteOnField(reflect.ValueOf(test), "someField", func(field reflect.StructField, value reflect.Value) {
		assert.Equal("someField", field.Name)
		assert.Equal("value", value.String())
		assert.Equal("thevalue", field.Tag.Get("sometag"))
	})
}

func TestForEachFieldTagged(t *testing.T) {
	assert := assert.New(t)

	test := testStruct{
		someField: "value",
	}

	ForEachFieldTagged(reflect.ValueOf(test), "sometag", func(field reflect.StructField, value reflect.Value, tagval string) {
		assert.Equal("someField", field.Name)
		assert.Equal("value", value.String())
		assert.Equal("thevalue", tagval)
	})
}

func TestForEachField(t *testing.T) {
	assert := assert.New(t)

	test := testStruct{
		someField: "value",
	}

	ForEachField(reflect.ValueOf(test), func(field reflect.StructField, value reflect.Value) {
		assert.Equal("someField", field.Name)
		assert.Equal("value", value.String())
	})
}

func TestFindFieldWithMatchingTag(t *testing.T) {
	assert := assert.New(t)

	test := testStruct{
		someField: "data",
	}
	field := FindFieldWithMatchingTag(reflect.ValueOf(test), "sometag", "thevalue")

	assert.Equal("data", field.String(), "field value is not data")
}

func TestIsEitherKind(t *testing.T) {
	assert := assert.New(t)

	testPtr := &testStruct{
		someField: "data",
	}

	test := testStruct{
		someField: "data",
	}

	assert.False(IsEitherKind(reflect.ValueOf(test), reflect.Pointer, reflect.Interface), "non pointer value is thought to be pointer")
	assert.True(IsEitherKind(reflect.ValueOf(testPtr), reflect.Pointer, reflect.Interface), "pointer value is thought to be non-pointer")
}

func TestResolveEditableValue(t *testing.T) {
	assert := assert.New(t)

	tInterface := makeInterface()
	testPtr := &testStruct{}

	assert.False(reflect.ValueOf(tInterface).CanSet())
	assert.False(reflect.ValueOf(testPtr).CanSet())
	assert.True(ResolveEditableValue(reflect.ValueOf(testPtr)).CanSet())
	assert.True(ResolveEditableValue(reflect.ValueOf(tInterface)).CanSet())
}

func makeInterface() testInterface {
	return &testStruct{}
}
