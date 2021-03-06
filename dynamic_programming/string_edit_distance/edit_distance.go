package main

import (
	"fmt"
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
	INSERT = 1
	DELETE = 2
)

type Cell struct {
	Operation int
	Cost      int
}

/*
	The cells in the matrix contains cost of the optimal solution to a sub problem, as well as
	parent pointer explaining how we got to this location.
*/
func createRecurrentRelationMatrix(x int, y int) [][]Cell {
	matrix := make([][]Cell, x+1)
	for i := 0; i <= x; i++ {
		matrix[uint8(i)] = make([]Cell, y+1)
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

func goalCell(str string, text string) (i int, j int) {
	return len(str), len(text)
}

func match(c string, d string) int {
	if c == d {
		return 0
	}

	return  1
}

func inDel(c string) int {
	return 1
}

func rowInt(i int, costMatrix *[][]Cell) {
	(*costMatrix)[0][i].Cost = i

	if i > 0 {
		(*costMatrix)[0][i].Operation = INSERT
	} else {
		(*costMatrix)[0][i].Operation = -1
	}
}

func columnInt(i int, costMatrix *[][]Cell) {
	(*costMatrix)[i][0].Cost = i

	if i > 0 {
		(*costMatrix)[i][0].Operation = DELETE
	} else {
		(*costMatrix)[i][0].Operation = -1
	}
}

func matchOut(str string, text string, i int, j int) {
	strChar := string(str[i - 1])
	textChar := string(text[j - 1])

	if strChar == textChar {
		fmt.Printf("M(%s, %s)->", strChar, textChar)
	} else  {
		fmt.Printf("S(%s, %s)->", strChar, textChar)
	}
}

func insertOut(text string, j int) {
	textChar := string(text[j-1])
	fmt.Printf("I(%s)->", textChar)
}

func deleteOut(str string, i int) {
	strChar := string(str[i-1])
	fmt.Printf("D(%s)->", strChar)
}

func reconstructPath(costMatrix [][]Cell, str string, text string, i int, j int) {
	if costMatrix[i][j].Operation == -1 {
		return
	}

	if costMatrix[i][j].Operation == MATCH {
		reconstructPath(costMatrix, str, text, i-1, j-1)
		matchOut(str, text, i, j)
		return
	}

	if costMatrix[i][j].Operation == INSERT {
		reconstructPath(costMatrix, str, text, i, j-1)
		insertOut(text, j)
		return
	}

	if costMatrix[i][j].Operation == DELETE {
		reconstructPath(costMatrix, str, text, i-1, j)
		deleteOut(str, i)
		return
	}
}

/*
	Input Examples:

	str := "thou shalt not murder"
	text := "you should not kill"

	str := "you should not"
	text := "thou shalt not"

	str := "victor"
	text := "viktor"
*/

func main() {
	str := "you should not"
	text := "thou shalt not"

	x := len(str)
	y := len(text)

	costMatrix := createRecurrentRelationMatrix(x, y)

	for i := 0; i <= y; i++ {
		rowInt(i, &costMatrix)
	}

	for i := 0; i <= x; i++ {
		columnInt(i, &costMatrix)
	}

	opt := make([]int, 3, 3)

	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {

			strChar := string(str[i-1])
			textChar := string(text[j-1])

			opt[MATCH] = costMatrix[i-1][j-1].Cost + match(strChar, textChar)
			opt[INSERT] = costMatrix[i][j-1].Cost + inDel(textChar)
			opt[DELETE] = costMatrix[i-1][j].Cost + inDel(strChar)

			costMatrix[i][j].Cost = opt[MATCH]
			costMatrix[i][j].Operation = MATCH

			for k := INSERT; k<=DELETE; k++ {
				if opt[k] < costMatrix[i][j].Cost {
					costMatrix[i][j].Cost = opt[k]
					costMatrix[i][j].Operation = k
				}
			}
		}
	}

	printRecurrentRelationMatrix(costMatrix)
	i, j := goalCell(str, text)

	fmt.Println("Edit Cost = ", costMatrix[i][j].Cost)

	reconstructPath(costMatrix, str, text, x, y)
}