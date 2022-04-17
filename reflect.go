package karen

import "reflect"

func FindFieldWithMatchingTag(object reflect.Value, key string, value string) reflect.Value {
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
