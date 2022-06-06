package karen

type Consumer[T any] func(value T)
type BiConsumer[T any, S any] func(first T, second S)
type Mapper[T any, R any] func(first T) R
type Filter[T any] func(value T) bool
