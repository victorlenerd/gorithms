// Leet Code: https://leetcode.com/problems/arithmetic-slices

package main

import "fmt"

func main()  {

	input := []int{1, 3, 5, 7, 9, 15, 20, 25, 28, 29}

	/*
		Creates a table the size of input to keep track of the number of arithmetic slice upto each item in the input
		and initializes them to 0
	*/
	dp := make([]int, len(input))
	sum := 0

	for i:=2; i<len(input); i++ {

		// Recurrent Relation
		if input[i] - input[i-1] == input[i-1] - input[i-2] {
			dp[i] = dp[i-1] + 1
			sum += dp[i]
		}
	}

	fmt.Println(dp, sum)
}
