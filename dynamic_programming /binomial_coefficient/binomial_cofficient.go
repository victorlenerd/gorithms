package main

import "fmt"

func createRecurrentRelationMatrix(n int) [][]int {
	size := n + 1
	matrix := make([][]int, size)
	for i:=0; i<=n; i++ {
		matrix[uint8(i)] = make([]int, size)
	}
	return matrix
}

func printRecurrentRelationMatrix(matrix [][]int )  {
	fmt.Println("---------------MATRIX------------------------")
	for i:=0; i<len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("---------------------------------------------")
}


func binomialCoefficient(n int, k int) {
	matrix := createRecurrentRelationMatrix(n)
	for i:=0; i<=n; i++ {
		for j:=0; j<=n; j++ {
			matrix[i][j] = 0
		}
	}
	// 1. base case: the number of ways to choose 0 things out of m possibilities is 1, an empty set.
	for m:= 0; m<=n; m++ {
		matrix[m][0] = 1
	}
	// 2. base case: the number of ways to choose m things out of m possibilities is 1, the complete set.
	for m:= 0; m<=n; m++ {
		matrix[m][m] = 1
	}
	printRecurrentRelationMatrix(matrix)
	for i := 1; i <= n; i++ {
		for j:=1; j<=i; j++ {
			matrix[i][j] = matrix[i-1][j-1] + matrix[i-1][j]
		}
	}
	printRecurrentRelationMatrix(matrix)
	fmt.Println("Binomial Coefficient:", matrix[n][k])
}

func main()  {
	n := 5
	k := 4
	binomialCoefficient(n, k)
}