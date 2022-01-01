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
	cost := []int{1,100,1,1,1,100,1,1,100,1}
	minimum := makeRecurrentRelationArr(cost)

	j := 2

	for j < len(minimum) {
		minimum[j] = int(math.Min(float64(minimum[j - 1] + cost[j - 1]),  float64(minimum[j - 2] + cost[j - 2])))
		j++
	}

	fmt.Println(minimum[len(minimum) - 1])
}
