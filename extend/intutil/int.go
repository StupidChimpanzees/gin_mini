package intutil

import (
	"sort"
	"strconv"
)

func IntToString(arr []int) []string {
	var strArr []string
	for _, v := range arr {
		strArr = append(strArr, strconv.Itoa(v))
	}
	return strArr
}

func InArray(arr *[]int, target int) bool {
	sort.Ints(*arr)
	index := sort.SearchInts(*arr, target)
	if index < len(*arr) && (*arr)[index] == target {
		return true
	}
	return false
}

func In8Array(arr *[]int8, target int8) bool {
	for _, u := range *arr {
		if target == u {
			return true
		}
	}
	return false
}
