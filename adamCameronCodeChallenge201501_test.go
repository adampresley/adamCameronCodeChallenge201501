package main

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestChallenge(t *testing.T) {
	var set = map[string]ChallengeItem{
		"one": {
			Alphabet: "a",
			Maori:    "tahi",
			Roman:    "i",
			Ordinal:  "first",
		},

		"two": {
			Alphabet: "b",
			Maori:    "rua",
			Roman:    "ii",
			Ordinal:  "second",
		},
		"three": {
			Alphabet: "c",
			Maori:    "toru",
			Roman:    "iii",
			Ordinal:  "third",
		},
		"four": {
			Alphabet: "d",
			Maori:    "wha",
			Roman:    "iv",
			Ordinal:  "fourth",
		},
	}

	Convey("Given a valid set of keys", t, func() {
		var keys = []string{
			"four", "two",
		}

		Convey("A subset is returned ordered by key array", func() {
			var expected = []ChallengeItem{
				{
					Alphabet: "d",
					Maori:    "wha",
					Roman:    "iv",
					Ordinal:  "fourth",
					Label:    "four",
				},
				{
					Alphabet: "b",
					Maori:    "rua",
					Roman:    "ii",
					Ordinal:  "second",
					Label:    "two",
				},
			}

			actual, _ := getOrderedSubset(set, keys)
			So(actual, ShouldResemble, expected)
		})

		Convey("Subset is in the correct order", func() {
			actual, _ := getOrderedSubset(set, keys)

			So(actual[0].Label, ShouldEqual, "four")
			So(actual[1].Label, ShouldEqual, "two")
		})
	})

	Convey("Given an invalid set of keys", t, func() {
		var keys = []string{
			"eight", "two",
		}

		Convey("An error will be returned", func() {
			expected := errors.New("Invalid key")
			_, err := getOrderedSubset(set, keys)
			So(err, ShouldResemble, expected)
		})
	})
}
