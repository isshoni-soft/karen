package karen

import (
	"container/list"
	"fmt"
	"reflect"
)

type Stack[T any] interface {
	Collection[T]

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

func (s *stack[T]) Add(value T) {
	s.Push(value)
}

func (s *stack[T]) Remove(value T) (bool, error) {
	elem, err := s.findElement(func(val T) bool {
		return reflect.DeepEqual(value, val)
	})

	if elem == nil {
		return false, err
	}

	s.list.Remove(elem)
	return true, err
}

func (s *stack[T]) Clear() {
	s.list = list.New()
}

func (s *stack[T]) AddCollection(add Collection[T]) error {
	return add.ForEach(func(val T) {
		s.Add(val)
	})
}

func (s *stack[T]) AddSlice(slice []T) {
	for _, v := range slice {
		s.Add(v)
	}
}

func (s *stack[T]) ForEach(consumer Consumer[T]) error {
	_, err := s.Find(func(value T) bool {
		consumer(value)
		return false
	})

	return err
}

func (s *stack[T]) findElements(filter Filter[T]) ([]*list.Element, error) {
	var result []*list.Element
	current := s.list.Front()

	for current != nil {
		t, err := CastError[T](current.Value)

		if err != nil {
			return nil, err
		}

		if filter(t) {
			result = append(result, current)
		}

		current = current.Next()
	}

	return result, nil
}

func (s *stack[T]) findElement(filter Filter[T]) (*list.Element, error) {
	current := s.list.Front()

	for current != nil {
		t, err := CastError[T](current.Value)

		if err != nil {
			return nil, err
		}

		if filter(t) {
			return current, err
		}

		current = current.Next()
	}

	return nil, nil
}

func (s *stack[T]) Find(filter Filter[T]) (T, error) {
	elem, err := s.findElement(filter)

	return Cast[T](elem.Value), err
}

func (s *stack[T]) AllMatching(filter Filter[T]) ([]T, error) {
	elems, err := s.findElements(filter)
	casted := MapSlice(elems, func(e *list.Element) T {
		return Cast[T](e.Value)
	})

	return casted, err
}

func (s *stack[T]) AsSlice() ([]T, error) {
	var result []T

	err := s.ForEach(func(val T) {
		result = append(result, val)
	})

	return result, err
}

func (s *stack[T]) Contains(value T) (bool, error) {
	result, err := s.Find(func(current T) bool {
		return reflect.DeepEqual(current, value)
	})

	return IsZero(result), err
}

func (s *stack[T]) Pop() (T, error) {
	if s.Empty() {
		return Zero[T](), fmt.Errorf("pop error: empty stack")
	}

	element := s.list.Front()

	s.list.Remove(element)

	return CastError[T](element.Value)
}

func (s stack[T]) Peek() (T, error) {
	if s.Empty() {
		return Zero[T](), fmt.Errorf("peek error: empty stack")
	}

	return CastError[T](s.list.Front().Value)
}

func (s stack[T]) Size() int {
	return s.list.Len()
}

func (s stack[T]) Empty() bool {
	return s.Size() == 0
}

func NewStack[T any]() Stack[T] {
	return &stack[T]{
		list: list.New(),
	}
}
