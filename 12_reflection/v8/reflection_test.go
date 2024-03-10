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
			"struct with one string field",
			struct {
				Name string
			}{"Hespecial"},
			[]string{"Hespecial"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Hespecial", "London"},
			[]string{"Hespecial", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Hespecial", 21},
			[]string{"Hespecial"},
		},
		{
			"nested fields",
			Person{
				"Hespecial",
				Profile{21, "London"},
			},
			[]string{"Hespecial", "London"},
		},
		{
			"pointers to things",
			&Person{
				"Hespecial",
				Profile{21, "London"},
			},
			[]string{"Hespecial", "London"},
		},
		{
			"slices",
			[]Profile{
				{21, "London"},
				{22, "China"},
			},
			[]string{"London", "China"},
		},
		{
			"arrays",
			[2]Profile{
				{21, "London"},
				{22, "China"},
			},
			[]string{"London", "China"},
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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
