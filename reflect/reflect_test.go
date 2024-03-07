package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Hespecial"},
			[]string{"Hespecial"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Hespecial", "Wuhan"},
			[]string{"Hespecial", "Wuhan"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Hespecial", 18},
			[]string{"Hespecial"},
		},
		{
			"Nested fields",
			Person{
				"Hespecial",
				Profile{18, "Wuhan"},
			},
			[]string{"Hespecial", "Wuhan"},
		},
		{
			"Pointers to things",
			&Person{
				"Hespecial",
				Profile{18, "Wuhan"},
			},
			[]string{"Hespecial", "Wuhan"},
		},
		{
			"Slices",
			[]Profile{
				{18, "Wuhan"},
				{20, "Shanghai"},
			},
			[]string{"Wuhan", "Shanghai"},
		},
		{
			"Arrays",
			[2]Profile{
				{18, "Wuhan"},
				{20, "Shanghai"},
			},
			[]string{"Wuhan", "Shanghai"},
		},
		{
			"Maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
