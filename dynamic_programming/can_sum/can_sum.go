package main

import "fmt"

/*
	Examples target 7, [2, 3] True
	Examples target 7, [5, 3, 4, 7] True
	Examples target 7, [2, 4] False
	Examples target 8, [2, 3, 5] True
	Examples target 300, [7, 14] False
*/

func makeRecurrentRelationArr(size int) []bool {
	res := make([]bool, size+1)
	i := 0
	for i < len(res) {
		res[i] = false
		i ++
	}
	return res
}

func main()  {
	target := 8
	arr := []int{2, 3, 5}

	matrix := makeRecurrentRelationArr(target)

	matrix[0] = true

	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(matrix); j++ {
			item := arr[i]

			if j >= item && !matrix[j] {
				complement := j - item
				matrix[j] = matrix[complement]
			}
		}
	}

	fmt.Println(matrix[target])
}