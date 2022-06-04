package karen

import (
	"fmt"
	"reflect"
)

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

func MapSlice[T any, R any](slice []T, mapper Mapper[T, R]) []R {
	var result []R

	for _, v := range slice {
		result = append(result, mapper(v))
	}

	return result
}

func CastAll[T any](incoming []any) []T {
	return MapSlice(incoming, func(val any) T {
		return Cast[T](val)
	})
}

func Cast[T any](incoming any) T {
	if incoming == nil {
		return Zero[T]()
	}

	if v, ok := incoming.(T); ok {
		return v
	}

	return Zero[T]()
}

func CastError[T any](i any) (T, error) {
	if v, ok := i.(T); ok {
		return v, nil
	} else {
		return Zero[T](), fmt.Errorf("conversion error: datatype in stack doesn't match high level type")
	}
}
