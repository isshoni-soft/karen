package karen

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type testInterface interface {
	Test() string
}

type testStruct struct {
	SomeField string `sometag:"thevalue"`
}

func (t testStruct) Test() string {
	return t.SomeField
}

func TestEditableValueOf(t *testing.T) {
	assert := assert.New(t)

	test := testStruct{
		SomeField: "value",
	}

	editableVal := EditableValueOf(&test)
	editableVal.FieldByName("SomeField").Set(reflect.ValueOf("test"))

	assert.Equal("test", test.SomeField)
}

func TestExecuteOnField(t *testing.T) {
	assert := assert.New(t)

	test := testStruct{
		SomeField: "value",
	}

	ExecuteOnField(reflect.ValueOf(test), "SomeField", func(field reflect.StructField, value reflect.Value) {
		assert.Equal("SomeField", field.Name)
		assert.Equal("value", value.String())
		assert.Equal("thevalue", field.Tag.Get("sometag"))
	})
}

func TestForEachFieldTagged(t *testing.T) {
	assert := assert.New(t)

	test := testStruct{
		SomeField: "value",
	}

	ForEachFieldTagged(reflect.ValueOf(test), "sometag", func(field reflect.StructField, value reflect.Value, tagval string) {
		assert.Equal("SomeField", field.Name)
		assert.Equal("value", value.String())
		assert.Equal("thevalue", tagval)
	})
}

func TestForEachField(t *testing.T) {
	assert := assert.New(t)

	test := testStruct{
		SomeField: "value",
	}

	ForEachField(reflect.ValueOf(test), func(field reflect.StructField, value reflect.Value) {
		assert.Equal("SomeField", field.Name)
		assert.Equal("value", value.String())
	})
}

func TestFindFieldWithMatchingTag(t *testing.T) {
	assert := assert.New(t)

	test := testStruct{
		SomeField: "data",
	}
	field := FindFieldWithMatchingTag(reflect.ValueOf(test), "sometag", "thevalue")

	assert.Equal("data", field.String(), "field value is not data")
}

func TestIsEitherKind(t *testing.T) {
	assert := assert.New(t)

	testPtr := &testStruct{
		SomeField: "data",
	}

	test := testStruct{
		SomeField: "data",
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
