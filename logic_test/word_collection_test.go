package logic_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordCollection(t *testing.T) {
	result := wordCollection([]string{"cook", "save", "taste", "aves", "vase", "state", "map"})

	expected := [][]string{
		{"cook"},
		{"save", "aves", "vase"},
		{"taste", "state"},
		{"map"},
	}

	assert.Equal(t, expected, result)
}

func wordCollection(words []string) [][]string {
	var results [][]string
	isAdded := make([]bool, len(words))

	for i := 0; i < len(words); i++ {
		if isAdded[i] {
			continue
		}

		var collection []string
		collection = append(collection, words[i])

		for j := i + 1; j < len(words); j++ {
			if isAdded[j] {
				continue
			}

			wordI := []rune(words[i])
			wordJ := []rune(words[j])

			var countedWordI rune
			var countedWordJ rune

			for _, item := range wordI {
				countedWordI += item
			}

			for _, item := range wordJ {
				countedWordJ += item
			}

			if countedWordI == countedWordJ {
				collection = append(collection, words[j])
				isAdded[j] = true
			}
		}

		results = append(results, collection)
	}

	return results
}
