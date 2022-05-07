package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) []string {
	keys := make([]int, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	result := make([]string, 0, len(keys))
	for _, v := range keys {
		result = append(result, input[v])
	}

	return result
}
