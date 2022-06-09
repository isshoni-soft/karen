package karen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type CastTest struct {
	data string
}

type FirstInterface interface {
	First()
}

type SecondInterface interface {
	Second()
}

func (c CastTest) First() {}

func (c CastTest) Second() {}

func TestZero(t *testing.T) {
	var str string

	if str != Zero[string]() {
		t.Fatal("zero does not match default string")
	}
}

func TestCast(t *testing.T) {
	test := CastTest{}

	if v := Cast[FirstInterface](test); v == nil {
		t.Fatal("CastTest unable to be cast to FirstInterface")
	}

	if v := Cast[SecondInterface](test); v == nil {
		t.Fatal("CastTest unable to be cast to SecondInterface")
	}

	if v := Cast[string](test); v != Zero[string]() {
		t.Fatal()
	}
}

func TestIsZero(t *testing.T) {
	assert := assert.New(t)

	test := Zero[CastTest]()

	assert.True(IsZero(test))
	assert.False(IsZero(CastTest{
		data: "test",
	}))
}

func TestIsType(t *testing.T) {
	assert := assert.New(t)
	test := CastTest{}

	assert.True(IsType[SecondInterface](test))
}
