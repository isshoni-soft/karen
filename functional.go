package karen

type Consumer[T any] func(value T)
type Mapper[T any, R any] func(first T) R
type Filter[T any] Mapper[T, bool]
