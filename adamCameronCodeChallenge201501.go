package main

import (
	"errors"
	"log"
)

/*
Represents a structure containing the values
specified in the challenge.
*/
type ChallengeItem struct {
	Alphabet string
	Maori    string
	Roman    string
	Ordinal  string
	Label    string
}

func main() {
	log.Println("INFO - Starting...")

	/*
	 * Make a map fo ChallengeItem items that looks like
	 * Mr. Cameron's example.
	 */
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

	/*
	 * Here are the keys we want to pull, and the
	 * order we want them in
	 */
	var keys = []string{"four", "two"}

	/*
	 * Call our function get our ordered ChallengeItem items
	 */
	result, err := getOrderedSubset(set, keys)
	if err != nil {
		log.Println("ERROR -", err)
	} else {
		log.Println(result)
	}
}

/*
Takes a set of ChallengeItem structs and an array of keys. This will return
an array of ChallengeItem structs for each key in the keys array. The resulting
ChallengeItem array will be ordered by the key order in the keys array.
*/
func getOrderedSubset(set map[string]ChallengeItem, keys []string) ([]ChallengeItem, error) {
	var result = make([]ChallengeItem, 0)

	for _, key := range keys {
		if challengeItem, ok := set[key]; ok {
			challengeItem.Label = key
			result = append(result, challengeItem)
		} else {
			return result, errors.New("Invalid key")
		}
	}

	return result, nil
}
