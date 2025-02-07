package uintutil

import (
	"strconv"
)

func UintToString(arr []uint) []string {
	var strArr []string
	for _, v := range arr {
		strArr = append(strArr, strconv.Itoa(int(v)))
	}
	return strArr
}

func Uint8ToString(arr []uint8) []string {
	var strArr []string
	for _, v := range arr {
		strArr = append(strArr, strconv.Itoa(int(v)))
	}
	return strArr
}

func InArray(arr *[]uint, target uint) bool {
	for _, u := range *arr {
		if target == u {
			return true
		}
	}
	return false
}

func In8Array(arr *[]uint8, target uint8) bool {
	for _, u := range *arr {
		if target == u {
			return true
		}
	}
	return false
}
