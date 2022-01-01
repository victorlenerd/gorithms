package main

import "fmt"

/**

To deal with in exact string matching we must first define a cost function telling us how far apart
two strings are i.e a distance measure between two pairs of strings.

The reasonable distance measure reflects the number of changes that must be made to convert one string to another.
There are three natural types of changes:

1. Substitution
2. Insertion
3. Deletions

Let i and j be the last character of the relevant prefix of P and T, respectively. There are three pair of shorter
strings after the last operation, corresponding to the strings after a match/substitution, insertion and deletion.

If we know the cost of editing these three pair of smaller strings, we could decide which option is best.

Let D[i, j] be the minimum number of differences between P1, P2,...Pi and the segment of T ending at j.
D[i, J] is the minimum of these options:

1. If (Pi == Tj), then D[i-1, j-1], else D[i-1, j-1] + 1. The means we either match or substitute the ith and jth
characters, depending upon whether the tail characters are the same.

2. D[i, j-1] + 1. This means that there is an extra character in the pattern to account for, so we do not advance the
next pointer and pay the cost of a deletion.

3. D[i - 1, j] + 1. This is means that there is an extra character in the text to remove, so we do not advance the pattern
pointer and pay the cost of insertion.
*/

const (
	MATCH  = 0
	INSERT = 1
	DELETE = 2
)

/*
	The cells in the matrix contains cost of the optimal solution to a sub problem, as well as
	parent pointer explaining how we got to this location.
*/
func createRecurrentRelationMatrix(m int, n int) [][]int {
	matrix := make([][]int, m + 1)
	for i:=0; i<m + 1; i++ {
		matrix[uint8(i)] = make([]int, n + 1)
	}
	return matrix
}

func printRecurrentRelationMatrix(matrix [][]int)  {
	fmt.Println("---------------MATRIX------------------------")
	for i:=0; i<len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("---------------------------------------------")
}

func findMin(x int, y int, z int) int {
	min := x

	if y < min {
		min = y
	}

	if z < min {
		min = z
	}

	return min
}

/*
	Input Examples:

	str := "thou shalt not murder"
	ptr := "you should not kill"

	str := "you should not"
	ptr := "thou shalt not"

	str := "victor"
	ptr := "viktor"
*/

func main() {
	str := "victor"
	ptr := "viktor"

	m := len(str)
	n := len(ptr)

	costMatrix := createRecurrentRelationMatrix(m, n)

	// Base Case: Giving an empty string and a pattern of i'th length the edit cost is i insertion
	for i := 1; i < m + 1; i++ {
		costMatrix[i][0] = i
	}

	// Base Case: Giving an empty pattern and a string of i'th length the edit cost is i insertion
	for i := 1; i < n + 1; i++ {
		costMatrix[0][i] = i
	}

	printRecurrentRelationMatrix(costMatrix)

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			strChar := string(str[i - 1])
			ptrChar := string(ptr[j - 1])

			// assume string and pattern char is a match then there's no substitution cost
			substitutionCost := costMatrix[i - 1][j - 1]

			// if string and pattern char do not match then substitution cost is factored
			if strChar != ptrChar {
				substitutionCost = costMatrix[i-1][j-1] + 1
			}

			costMatrix[i][j] = findMin(
				substitutionCost, // Cost of substitution
				costMatrix[i][j-1] + 1, // Cost of deletion
				costMatrix[i-1][j] + 1, // Cost of insertion
			)
		}
	}

	printRecurrentRelationMatrix(costMatrix)

	fmt.Println("Edit Cost", costMatrix[m][n])
}
