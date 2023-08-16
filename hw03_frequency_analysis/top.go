package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

func Top10(inputString string) (output []string) {
	rgxp := regexp.MustCompile(`\s{1,}`)
	inputString = rgxp.ReplaceAllString(inputString, " ")

	dict := make(map[string]int)
	output = make([]string, 0)
	count := 0
	const TopCount = 10

	r := regexp.MustCompile(`[^\s]+`)
	wordsSlice := r.FindAllString(inputString, -1)

	for _, value := range wordsSlice {
		dict[value]++
	}

	keys := make([]string, 0, len(dict))

	for key := range dict {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		if dict[keys[i]] > dict[keys[j]] {
			return true
		} else if dict[keys[i]] == dict[keys[j]] {
			compareResult := strings.Compare(keys[i], keys[j])
			return compareResult == -1
		}
		return false
	})

	for _, k := range keys {
		count++
		if count > TopCount {
			break
		}
		output = append(output, k)
	}

	return
}
