package maplookuppattern

func HasDuplicate(data []int) bool {
	seen := make(map[int]bool)
	for _, v := range data {
		if _, ok := seen[v]; ok {
			return true
		}
		seen[v] = true
	}

	return false
}

func IsString(val interface{}) bool {
	if _, ok := val.(string); ok {
		return true
	}

	return false
}
