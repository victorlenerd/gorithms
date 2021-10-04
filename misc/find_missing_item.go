package misc

func FindMissingItem(a []int, b []int) []int {
	var missing []int

	aMap := make(map[int]int)

	for _, ac := range a {
		aMap[ac] = ac
	}

	for _, bc := range b {
		if _, ok := aMap[bc]; ok {
			delete(aMap, bc)
		}
	}

	for _, v := range aMap {
		missing = append(missing, v)
	}

	return missing
}
