package main

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	v := reflect.ValueOf(x)
	fn(v.Field(0).String())
}
