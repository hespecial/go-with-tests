package main

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	v := getVal(x)

	walkValue := func(val reflect.Value) {
		walk(val.Interface(), fn)
	}

	switch v.Kind() {
	case reflect.String:
		fn(v.String())
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			walkValue(v.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			walkValue(v.Index(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			walkValue(v.MapIndex(key))
		}
	}
}

func getVal(x interface{}) reflect.Value {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	return v
}
