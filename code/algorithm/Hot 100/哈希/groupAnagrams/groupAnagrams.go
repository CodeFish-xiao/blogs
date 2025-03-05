package groupAnagrams

import (
	"maps"
	"slices"
)

// 异构分组
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}
	m := make(map[string][]string)
	for _, str := range strs {
		s := []byte(str)
		slices.Sort(s)
		m[string(s)] = append(m[string(s)], str)

	}
	return slices.Collect(maps.Values(m))
}
