package topwords

import (
	"sort"
	"strings"
	"unicode"
)

func Top10(s string) []string {

	var result []string

	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	slice := strings.FieldsFunc(s, f)

	cache := make(map[string]int, len(slice))

	var i, j, c int
	for i = 0; i < len(slice); i++ {
		c = 0
		for j = i; j < len(slice); j++ {
			if slice[i] == slice[j] {
				c++
			}
		}

		_, ok := cache[slice[i]]
		if ok == true {

			continue

		}

		if c > 1 {
			cache[slice[i]] = c
		}

	}

	type keyValue struct {
		Key   string
		Value int
	}

	var sortedStruct []keyValue

	for key, value := range cache {
		sortedStruct = append(sortedStruct, keyValue{key, value})
	}

	sort.Slice(sortedStruct, func(i, j int) bool {
		return sortedStruct[i].Value > sortedStruct[j].Value
	})

	if len(sortedStruct) > 10 {
		sortedStruct = sortedStruct[:10]
	}

	for _, keyValue := range sortedStruct {
		result = append(result, keyValue.Key)

	}
	return result
}
