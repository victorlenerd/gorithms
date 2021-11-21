package main

import (
	"fmt"
	"math"
)

func makeRecurrentRelationArr(arr []int) []int {
	res := make([]int, len(arr) +1)
	i := 0
	for i < len(res) {
		res[i] = 0
		i ++
	}
	return res
}


func main()  {
	input := []int{2, 1, 4, 5, 8}

	size := len(input)

	maximums := makeRecurrentRelationArr(input)

	maximums[size] = 0
	maximums[size - 1] = input[size - 1]

	i := size - 2

	for i >= 0 {
		maximums[i] = int(math.Max(float64(maximums[i + 1]), float64(maximums[i + 2] + input[i])))
		i = i - 1
	}

	fmt.Println(maximums)
}
