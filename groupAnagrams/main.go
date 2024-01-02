package groupAnagrams

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	anagramsGrouped := make(map[string][]string)
	for _, element := range strs {
		slice := strings.Split(element, "")
		sort.Strings(slice)
		sorted := strings.Join(slice, "")
		anagramsGrouped[sorted] = append(anagramsGrouped[sorted], element)
	}

	result := make([][]string, 0, len(anagramsGrouped))

	for _, g := range anagramsGrouped {
		result = append(result, g)
	}

	return result
}
