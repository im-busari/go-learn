package reflection

import "reflect"

// golang challenge: write a function walk(x interface{}, fn func(string)) which takes a struct x and calls
// fn for all strings fields found inside. difficulty level: recursively.

// Reflections do checks at runtime

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		if field.Kind() == reflect.String {
			fn(field.String())
		}

		// To handle nested fields
		if field.Kind() == reflect.Struct {
			walk(field.Interface(), fn)
		}
	}
}
