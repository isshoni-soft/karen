package karen

type Set[T comparable] interface {
	Collection[T]
}

type set[T comparable] struct {
	Set[T]

	backing map[T]any
}

func (s *set[T]) Add(value T) {
	s.backing[value] = nil
}

func (s *set[T]) Remove(value T) (bool, error) {
	if _, ok := s.backing[value]; !ok {
		return false, nil
	}

	delete(s.backing, value)

	return true, nil
}

func (s *set[T]) AddCollection(add Collection[T]) error {
	return add.ForEach(func(v T, _ int) {
		s.Add(v)
	})
}

func (s *set[T]) AddSlice(add []T) {
	for _, v := range add {
		s.Add(v)
	}
}

func (s *set[T]) AllMatching(filter Filter[T]) ([]T, error) {
	var result []T

	for k := range s.backing {
		if filter(k) {
			result = append(result, k)
		}
	}

	return result, nil
}

func (s *set[T]) Find(filter Filter[T]) (T, error) {
	for k := range s.backing {
		if filter(k) {
			return k, nil
		}
	}

	return Zero[T](), nil
}

func (s *set[T]) ForEach(consumer BiConsumer[T, int]) error {
	index := 0

	_, err := s.Find(func(current T) bool {
		consumer(current, index)
		index++
		return false
	})

	return err
}

func (s *set[T]) Clear() {
	s.backing = make(map[T]any)
}

func (s *set[T]) AsSlice() ([]T, error) {
	var result []T

	for k := range s.backing {
		result = append(result, k)
	}

	return result, nil
}

func (s *set[T]) Contains(value T) (bool, error) {
	for k := range s.backing {
		if value == k {
			return true, nil
		}
	}

	return false, nil
}

func (s *set[T]) Size() int {
	return len(s.backing)
}

func (s *set[T]) Empty() bool {
	return s.Size() == 0
}

func NewSet[T comparable]() Set[T] {
	return &set[T]{
		backing: make(map[T]any),
	}
}
