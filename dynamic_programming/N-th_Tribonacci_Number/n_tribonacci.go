package main

import "fmt"

func main() {
	memo := make(map[int]int)
	n := 3
	fmt.Println("tribonacci", tribonacci(n, &memo))
}

func tribonacci(n int, memo *map[int]int) int {
	if trib, ok := (*memo)[n]; ok {
		return trib
	}

	if n <= 0 {
		return 0
	}

	if n <= 2 {
		return 1
	}

	(*memo)[n] = tribonacci(n - 1, memo) + tribonacci(n - 2, memo) + tribonacci(n - 3, memo)

	return (*memo)[n]
}

