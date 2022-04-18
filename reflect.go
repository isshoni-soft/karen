package karen

import "reflect"

func ResolveEditableValue(object reflect.Value) reflect.Value {
	if IsEitherKind(object, reflect.Interface, reflect.Pointer) {
		object = object.Elem()
		return ResolveEditableValue(object)
	}

	return object
}

func IsEitherKind(object reflect.Value, kinds ...reflect.Kind) bool {
	for _, k := range kinds {
		if object.Kind() == k {
			return true
		}
	}

	return false
}

func FindFieldWithMatchingTag(object reflect.Value, key string, value string) reflect.Value {
	if IsEitherKind(object, reflect.Interface, reflect.Pointer) {
		object = object.Elem()
	}

	objectType := object.Type()
	numFields := objectType.NumField()

	for x := 0; x < numFields; x++ {
		field := objectType.Field(x)

		if v, ok := field.Tag.Lookup(key); ok && v == value {
			return object.Field(x)
		}
	}

	return reflect.Value{}
}
