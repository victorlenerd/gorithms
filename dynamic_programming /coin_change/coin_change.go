package main

import "fmt"

func makeRecurrentRelationArr(size int) []int {
	res := make([]int, size+1)
	i := 0
	for i < len(res) {
		res[i] = 0
		i ++
	}
	return res
}

func main() {

	coins := []int{2, 5, 3, 6}
	targetMoney := 10
	change := makeRecurrentRelationArr(targetMoney)

	// Base case: for coin 0 the minimum number of change is 0
	change[0] = 1

	for _, currentCoin := range coins {
		money := currentCoin

		for money <= targetMoney {
			diff := money - currentCoin

			if change[diff] > 0 {

				// for change[x] there are N additional ways to change it
				change[money] += change[diff]
			}

			money++
		}


		fmt.Println("change", change)

	}

}
