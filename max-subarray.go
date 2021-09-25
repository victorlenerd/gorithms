package main

import (
	"math"
)

func maxSubarray(arr []float64) float64 {
	globalMax := arr[0]
	localMax := globalMax

	for _, val := range arr[1:] {
		localMax = math.Max(val, localMax+val)

		if localMax > globalMax {
			globalMax = localMax
		}
	}

	return globalMax
}
