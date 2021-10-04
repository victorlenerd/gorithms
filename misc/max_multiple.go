package misc

import "fmt"

const BASE = 5
const CELINE = 300

func MaxMultiple(pos int, currentFactorial int) int {
	if pos > CELINE {
		return BASE
	}

	if currentFactorial >= CELINE {
		return currentFactorial
	}

	if currentFactorial >= pos {
		return currentFactorial
	}

	return MaxMultiple(pos, currentFactorial+BASE)
}

func main() {
	pos := 7
	currentFactorial := 0

	max := MaxMultiple(pos, currentFactorial)

	start := 0
	stop := max

	if max == pos {
		start = pos
		stop = max + BASE
	} else {
		start = max - BASE
		stop = max
	}

	for i := start; i < stop; i++ {
		if i == pos {
			fmt.Println(i, "-->")
		} else {
			fmt.Println(i)
		}
	}
}
