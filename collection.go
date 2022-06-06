package karen

type Collection[T any] interface {
	Add(value T)
	Remove(value T) (bool, error)
	AddCollection(collection Collection[T]) error
	AddSlice(slice []T)
	ForEach(consumer BiConsumer[T, int]) error
	Find(filter Filter[T]) (T, error)
	AllMatching(filter Filter[T]) ([]T, error)
	Contains(value T) (bool, error)
	Clear()
	AsSlice() ([]T, error)
	Size() int
	Empty() bool
}
