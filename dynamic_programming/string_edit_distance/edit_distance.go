package main

import (
	"fmt"
	"math"
)

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
	SUB    = 1
	INSERT = 2
	DELETE = 3
)

type Cell struct {
	Operation int
	Cost      int
}

/*
	The cells in the matrix contains cost of the optimal solution to a sub problem, as well as
	parent pointer explaining how we got to this location.
*/
func createRecurrentRelationMatrix(m int, n int) [][]Cell {
	matrix := make([][]Cell, m+1)
	for i := 0; i < m+1; i++ {
		matrix[uint8(i)] = make([]Cell, n+1)
	}
	return matrix
}

func printRecurrentRelationMatrix(matrix [][]Cell) {
	fmt.Println("---------------MATRIX------------------------")
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("---------------------------------------------")
}

func findMin(isMatch bool, match int, substitution int, insertion int, deletion int) (int, int) {
	min := math.MaxInt8
	var operation int

	if isMatch {
		min = match
		operation = MATCH
	} else {
		min = substitution
		operation = SUB
	}

	if insertion < min {
		min = insertion
		operation = INSERT
	}

	if deletion < min {
		min = deletion
		operation = DELETE
	}

	return min, operation
}

func printOperations(matrix [][]Cell, i int, j int)  {

	queue := []int{
		matrix[i][j].Operation,
	}

	for len(queue) > 0 {
		top := queue[0]

		if top == MATCH {
			fmt.Print("M ->")
			queue = append(queue, matrix[i-1][j-1].Operation)
			i -= 1
			j -= 1
		}

		if top == SUB {
			fmt.Print("S ->")
			queue = append(queue, matrix[i-1][j-1].Operation)
			i -= 1
			j -= 1
		}

		if top == INSERT {
			fmt.Print("I ->")
			queue = append(queue, matrix[i][j-1].Operation)
			j -= 1
		}

		if top == DELETE {
			fmt.Print("D ->")
			queue = append(queue, matrix[i-1][j].Operation)
			i -= 1
		}

		queue = queue[1:]
	}

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
	ptr := "vik"
	str := "victor"

	m := len(str)
	n := len(ptr)

	costMatrix := createRecurrentRelationMatrix(m, n)

	costMatrix[0][0] = Cell{
		Cost:      0,
		Operation: -1,
	}

	// Base Case: Giving an empty string and a pattern of i'th length the edit cost is i insertion
	for i := 1; i < m+1; i++ {
		costMatrix[i][0] = Cell{
			Cost:      i,
			Operation: INSERT,
		}
	}

	// Base Case: Giving an empty pattern and a string of i'th length the edit cost is i deletion
	for i := 1; i < n+1; i++ {
		costMatrix[0][i] = Cell{
			Cost:      i,
			Operation: DELETE,
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			strChar := string(str[i-1])
			ptrChar := string(ptr[j-1])

			min, operation := findMin(
				strChar == ptrChar,
				costMatrix[i-1][j-1].Cost,   // Cost or a match
				costMatrix[i-1][j-1].Cost+1, // Cost of a substitution
				costMatrix[i][j-1].Cost+1,   // Cost of an insertion
				costMatrix[i-1][j].Cost+1,   // Cost of a deletion
			)

			cell := Cell{
				Cost:      min,
				Operation: operation,
			}

			costMatrix[i][j] = cell
		}
	}

	printRecurrentRelationMatrix(costMatrix)

	fmt.Println("Edit Cost", costMatrix[m][n])
	printOperations(costMatrix, m, n)
}
