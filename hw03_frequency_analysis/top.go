package hw03frequencyanalysis

import (
	"regexp"
	"sort"
)

func Top10(inputString string) (output []string) {
	rgxp := regexp.MustCompile(`\s{1,}`)
	inputString = rgxp.ReplaceAllString(inputString, " ")

	dict := make(map[string]int)
	output = make([]string, 0)
	keys := make([]string, 0, len(dict))
	count := 0
	const TopCount = 10

	r := regexp.MustCompile(`[^\s]+`)
	wordsSlice := r.FindAllString(inputString, -1)

	for _, value := range wordsSlice {
		dict[value]++
	}

	for key := range dict {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	sort.SliceStable(keys, func(i, j int) bool {
		return dict[keys[i]] > dict[keys[j]]
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
