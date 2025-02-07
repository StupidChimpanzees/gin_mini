package stringutil

import "sort"

func InArray(arr *[]string, target string) bool {
	sort.Strings(*arr)
	index := sort.SearchStrings(*arr, target)
	if index < len(*arr) && (*arr)[index] == target {
		return true
	}
	return false
}
