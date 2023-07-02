package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	type words struct {
		word  string
		count int
	}
	s := strings.Fields(text)
	rawDict := make(map[string]int)
	for _, val := range s {
		rawDict[val] = 0
	}
	for i := 0; i < len(s); i++ {
		for key := range rawDict {
			if s[i] == key {
				rawDict[key]++
			}
		}
	}
	dictWords := make([]words, 0, len(rawDict))
	for k, v := range rawDict {
		dictWords = append(dictWords, words{k, v})
	}
	sort.Slice(dictWords, func(i, j int) bool {
		if dictWords[i].count == dictWords[j].count {
			return dictWords[i].word < dictWords[j].word
		}
		return dictWords[i].count > dictWords[j].count
	})
	result := make([]string, 0, 10)
	if len(dictWords) >= 10 {
		for i := 0; i < 10; i++ {
			result = append(result, dictWords[i].word)
		}
	} else {
		for _, val := range dictWords {
			result = append(result, val.word)
		}
	}
	return result
}
