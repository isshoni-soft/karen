package karen

func Zero[T any]() T {
	var result T
	return result
}

func Cast[T any](incoming any) T {
	if v, ok := incoming.(T); ok {
		return v
	}

	return Zero[T]()
}
