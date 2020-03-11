package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	t.Run("test walk without maps", func(t *testing.T) {
		cases := []struct {
			Name          string
			Input         interface{}
			ExpectedCalls []string
		}{
			{
				"Struct with one string field",
				struct {
					Name string
				}{"Chris"},
				[]string{"Chris"},
			},
			{
				"Struct with two string fields",
				struct {
					Name  string
					Email string
				}{"John", "john@email.com"},
				[]string{"John", "john@email.com"},
			},
			{
				"Struct with non string field",
				struct {
					Name string
					Age  int
				}{"Jess", 29},
				[]string{"Jess"},
			},
			{
				"Nested fields",
				Person{
					"Chris",
					Profile{33, "London"},
				},
				[]string{"Chris", "London"},
			},
			{
				"pointers to things",
				&Person{
					"Chris",
					Profile{22, "London"},
				},
				[]string{"Chris", "London"},
			},
			{
				"slices",
				[]Profile{
					{33, "London"},
					{34, "Reykjavik"},
				},
				[]string{"London", "Reykjavik"},
			},
			{
				"arrays",
				[2]Profile{
					{33, "London"},
					{34, "Reykjavik"},
				},
				[]string{"London", "Reykjavik"},
			},
			{
				"maps",
				map[int]string{
					1: "Dog",
					2: "Cat",
				},
				[]string{"Dog", "Cat"},
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
	})

	t.Run("test walk with maps", func(t *testing.T) {
		aMap := map[int]string{
			1: "Dog",
			2: "Cat",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Cat")
		assertContains(t, got, "Dog")
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
