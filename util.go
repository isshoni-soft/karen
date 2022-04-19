package karen

import "reflect"

func Zero[T any]() T {
	var result T
	return result
}

func IsZero[T any](incoming T) bool {
	zero := Zero[T]()

	return reflect.DeepEqual(incoming, zero)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Cast[T any](incoming any) T {
	if v, ok := incoming.(T); ok {
		return v
	}

	return Zero[T]()
}
