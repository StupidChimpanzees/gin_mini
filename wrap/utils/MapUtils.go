package utils

func MergeMaps(map1, map2 map[string]any) map[string]any {
	for k, v := range map2 {
		if map1[k] != nil {
			value1, ok1 := map1[k].(map[string]any)
			value2, ok2 := v.(map[string]any)
			if ok1 && ok2 {
				MergeMaps(value1, value2)
			} else {
				map1[k] = v
			}
		} else {
			map1[k] = v
		}
	}
	return map1
}
