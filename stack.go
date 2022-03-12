package karen

import (
	"container/list"
	"fmt"
)

type Stack[T any] interface {
	Push(value T)
	Pop() (T, error)
	Peek() (T, error)
	Size() int
	Empty() bool
}

type stack[T any] struct {
	Stack[T]

	list *list.List
}

func (s *stack[T]) Push(value T) {
	s.list.PushFront(value)
}

func (s *stack[T]) Pop() (T, error) {
	if s.Empty() {
		return Zero[T](), fmt.Errorf("pop error: empty stack")
	}

	element := s.list.Front()

	s.list.Remove(element)

	return s.castError(element.Value)
}

func (s stack[T]) Peek() (T, error) {
	if s.Empty() {
		return Zero[T](), fmt.Errorf("peek error: empty stack")
	}

	return s.castError(s.list.Front().Value)
}

func (s stack[T]) Size() int {
	return s.list.Len()
}

func (s stack[T]) Empty() bool {
	return s.Size() == 0
}

func (s stack[T]) castError(i any) (T, error) {
	if v, ok := i.(T); ok {
		return v, nil
	} else {
		return Zero[T](), fmt.Errorf("conversion error: datatype in stack doesn't match high level type")
	}
}

func NewStack[T any]() Stack[T] {
	return &stack[T]{
		list: list.New(),
	}
}
