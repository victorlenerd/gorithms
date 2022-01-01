package main

import "fmt"

func main()  {

	nums := []int{0,2,3}
	dp := make([]bool, len(nums))
	dp[0] = true

	for i := 1; i<len(nums); i++ {
		curr := nums[i-1]
		for k := 0; k < curr; k++ {
			if i+k <= len(nums) - 1 && dp[i-1] {
				if !dp[i+k] {
					dp[i+k] = true
				}
			}
		}
	}

	fmt.Println(dp)
}
