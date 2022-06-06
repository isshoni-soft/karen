package karen

import (
	"testing"
)

var expected = []string{"string3", "string2", "string1"}

func makeStack() Stack[string] {
	stack := NewStack[string]("string1", "string2", "string3")

	return stack
}

func validateSlice(t *testing.T, slice []string) {
	for i, s := range expected {
		if slice[i] != s {
			t.Fatal("produced string slice does not match expected", slice)
		}
	}
}

func checkError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestStackPeek(t *testing.T) {
	stack := makeStack()

	if v, err := stack.Peek(); err == nil && v != "string3" {
		t.Fatal("top of stack not 'string3'")
	} else {
		checkError(t, err)
	}
}

func TestStackEmpty(t *testing.T) {
	stack := NewStack[string]()

	if !stack.Empty() {
		t.Fatal("stack should be empty")
	}
}

func TestStackSize(t *testing.T) {
	stack := makeStack()

	if stack.Size() != 3 {
		t.Fatal("size should be 3")
	}
}

func TestStackPop(t *testing.T) {
	stack := makeStack()

	if v, err := stack.Pop(); err == nil && v != "string3" {
		t.Fatal("top of stack not 'string3'")
	} else {
		checkError(t, err)
	}
}

func TestStackNotContains(t *testing.T) {
	stack := makeStack()

	if ok, err := stack.Contains("bad"); ok {
		t.Fatal("stack says it contains 'bad' but it doesn't")
	} else {
		checkError(t, err)
	}
}

func TestStackContains(t *testing.T) {
	stack := makeStack()

	if ok, err := stack.Contains("string1"); !ok {
		t.Fatal("stack says it does not contain 'string1' but it doesn't")
	} else {
		checkError(t, err)
	}
}

func TestStackAsSlice(t *testing.T) {
	stack := makeStack()

	slice, err := stack.AsSlice()

	checkError(t, err)
	validateSlice(t, slice)
}

func TestStackAllMatching(t *testing.T) {
	stack := makeStack()

	matching, err := stack.AllMatching(func(value string) bool {
		return value[0] == 's'
	})

	checkError(t, err)

	if len(matching) != 3 {
		t.Fatal("did not find all strings")
	}

	validateSlice(t, matching)
}

func TestStackFind(t *testing.T) {
	stack := makeStack()

	found, err := stack.Find(func(value string) bool {
		return value[0] == 's'
	})

	checkError(t, err)

	if found != "string3" {
		t.Fatal("didn't find first item in the stack")
	}
}

func TestStackForEach(t *testing.T) {
	stack := makeStack()

	err := stack.ForEach(func(value string, index int) {
		if expected[index] != value {
			t.Fatal("for each value doesn't match index")
		}
	})

	checkError(t, err)
}
