package utils

func StrArrToMultiMap(array []string, value any) map[string]any {
	multiMap := make(map[string]any)
	sliceArr := array[1:]
	if len(sliceArr) > 0 {
		multiMap[array[0]] = StrArrToMultiMap(sliceArr, value)
	} else {
		multiMap[array[0]] = value
	}
	return multiMap
}

func StrUnique(m []string) []string {
	d := make([]string, 0)
	tempMap := make(map[string]bool, len(m))
	for _, v := range m {
		if tempMap[v] == false {
			tempMap[v] = true
			d = append(d, v)
		}
	}
	return d
}
