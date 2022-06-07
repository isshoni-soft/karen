package karen

import "testing"

func makeSet() Set[string] {
	return NewSet[string](expected...)
}

func validateSet(t *testing.T, slice, against []string) {
	if len(slice) != len(against) {
		t.Fatal("slices not equal")
	}

	var found []string

	for _, s := range slice {
		for _, s2 := range against {
			if s == s2 {
				found = append(found, s)
			}
		}
	}

	if len(found) != len(slice) {
		t.Fatal("slices are not equal")
	}
}

func TestSetEmpty(t *testing.T) {
	set := NewSet[string]()

	if !set.Empty() {
		t.Fatal("set reporting not empty")
	}
}

func TestSetAdd(t *testing.T) {
	set := NewSet[string]()

	set.Add("string1")

	if ok, err := set.Contains("string1"); !ok {
		t.Fatal("set reports it does not contain string1")
	} else {
		checkError(t, err)
	}
}

func TestSetRemove(t *testing.T) {
	set := makeSet()

	if removed, err := set.Remove("string1"); !removed {
		t.Fatal("set reports it did not remove string1")
	} else {
		checkError(t, err)
	}

	if set.Size() != 2 {
		t.Fatal("set size did not change")
	}
}

func TestSetAddCollection(t *testing.T) {
	set := NewSet[string]()
	stack := NewStack[string]("string1", "string2", "string3")

	err := set.AddCollection(stack)

	checkError(t, err)

	slice, err := set.AsSlice()

	checkError(t, err)
	validateSet(t, slice, expected)
}

func TestSetAddSlice(t *testing.T) {
	set := NewSet[string]()

	set.AddSlice(expected)

	slice, err := set.AsSlice()

	checkError(t, err)
	validateSet(t, slice, expected)
}

func TestSetAllMatching(t *testing.T) {
	set := makeSet()
	set.Add("test")

	matching, err := set.AllMatching(func(value string) bool {
		return value[0] == 's'
	})

	checkError(t, err)
	validateSet(t, matching, expected)
}

func TestSetFind(t *testing.T) {
	set := makeSet()

	found, err := set.Find(func(value string) bool {
		return value == "string2"
	})

	checkError(t, err)

	if found != "string2" {
		t.Fatal("found is not string2")
	}
}

func TestSetForEach(t *testing.T) {
	set := makeSet()

	var found []string

	err := set.ForEach(func(value string, index int) {
		found = append(found, value)
	})

	checkError(t, err)
	validateSet(t, found, expected)
}

func TestSetClear(t *testing.T) {
	set := makeSet()

	set.Clear()

	if !set.Empty() {
		t.Fatal("cleared set not reporting empty")
	}

	if set.Size() != 0 {
		t.Fatal("cleared set size is not 0")
	}
}

func TestSetSize(t *testing.T) {
	set := makeSet()

	if set.Size() != 3 {
		t.Fatal("set size is not 3")
	}
}

func TestSetAsSlice(t *testing.T) {
	set := makeSet()

	slice, err := set.AsSlice()

	checkError(t, err)
	validateSet(t, slice, expected)
}

func TestSetContains(t *testing.T) {
	set := makeSet()

	if ok, err := set.Contains("string2"); !ok {
		t.Fatal("set reports it does not contains string2")
	} else {
		checkError(t, err)
	}
}
