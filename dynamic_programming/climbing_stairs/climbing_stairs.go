// https://leetcode.com/problems/climbing-stairs/

package main

import "fmt"

func main() {
	fmt.Println("climbStairs", climbStairs(45))
}

func climbStairsWithMemo(n int, memo map[int]int) int {
	if r, ok := memo[n]; ok {
		return r
	}

	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	memo[n] = climbStairsWithMemo(n - 1, memo) + climbStairsWithMemo(n - 2, memo)

	return memo[n]
}

func climbStairs(n int) int {
	memo := map[int]int{}

	return climbStairsWithMemo(n - 1, memo) + climbStairsWithMemo(n - 2, memo)
}
