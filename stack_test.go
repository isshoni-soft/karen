package karen

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack[string]()

	if stack == nil {
		t.Fatal("failed to create new stack!")
	}
}

func TestStackBasics(t *testing.T) {
	stack := NewStack[string]()

	stack.Push("string1")
	stack.Push("string2")
	stack.Push("string3")

	if v, err := stack.Peek(); err == nil && v != "string3" {
		t.Fatal("top of stack not 'string3'")
	} else if err != nil {
		t.Fatal(err)
	}
}
