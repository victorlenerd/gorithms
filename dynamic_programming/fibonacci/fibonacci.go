package main

import "fmt"

func main() {
	memo := make(map[int]int)
	n := 3
	fmt.Println("fibonacci", fibonacci(n, &memo))
}

func fibonacci(n int, memo *map[int]int) int {
	if fib, ok := (*memo)[n]; ok {
		return fib
	}

	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	(*memo)[n] = fibonacci(n - 1, memo) + fibonacci(n - 2, memo)

	return (*memo)[n]
}