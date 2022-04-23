package karen

import "reflect"

func EditableValueOf(value any) reflect.Value {
	return ResolveEditableValue(reflect.ValueOf(value))
}

func InsertIntoField(object reflect.Value, fieldName string, value any) (any, bool) {
	success := false
	var prevVal any

	ExecuteOnField(object, fieldName, func(field reflect.StructField, val reflect.Value) {
		prevVal = val.Interface()

		if val.CanSet() {
			success = true
			val.Set(reflect.ValueOf(value))
		}
	})

	return prevVal, success
}

func ExecuteOnField(object reflect.Value, fieldName string, consumer func(field reflect.StructField, value reflect.Value)) {
	object = ResolveEditableValue(object)
	field, exists := object.Type().FieldByName(fieldName)
	if !exists {
		return
	}

	consumer(field, object.FieldByName(fieldName))
}

func ForEachFieldTagged(object reflect.Value, tagName string, consumer func(field reflect.StructField, value reflect.Value, tagValue string)) {
	ForEachField(object, func(field reflect.StructField, value reflect.Value) {
		if v, ok := field.Tag.Lookup(tagName); ok {
			consumer(field, value, v)
		}
	})
}

func ForEachField(object reflect.Value, consumer func(field reflect.StructField, value reflect.Value)) {
	object = ResolveEditableValue(object)
	fieldNums := object.NumField()

	for x := 0; x < fieldNums; x++ {
		field := object.Type().Field(x)
		consumer(field, object.Field(x))
	}
}

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
