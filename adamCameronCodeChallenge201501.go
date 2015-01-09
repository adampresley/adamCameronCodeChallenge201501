package main

import (
	"errors"
	"log"
)

type ChallengeItem struct {
	Alphabet string
	Maori    string
	Roman    string
	Ordinal  string
	Label    string
}

func main() {
	log.Println("INFO - Starting...")

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

	var keys = []string{"four", "two"}

	result, err := getOrderedSubset(set, keys)
	if err != nil {
		log.Println("ERROR -", err)
	} else {
		log.Println(result)
	}
}

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
