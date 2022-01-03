package main

import (
	"fmt"
)

const (
	MATCH  = 0
	INSERT = 1
	DELETE = 2
)

type Cell struct {
	Operation int
	Cost      int
}

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

func goalCell(costMatrix [][]Cell,str string, text string) (i int, j int) {
	var k int

	i = len(str)
	j = 0

	for k=1; k<len(text); k++ {
		if costMatrix[i][k].Cost < costMatrix[i][j].Cost {
			j = k
		}
	}

	return i, j
}

func match(c string, d string) int {
	if c == d {
		return 0
	}

	return  1
}


func insert() int {
	return 1
}

func del() int {
	return 1
}

func rowInt(j int, costMatrix *[][]Cell) {
	(*costMatrix)[0][j].Cost = 0
	(*costMatrix)[0][j].Operation = -1
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
	str := "vic"
	text := "sstorpu3ddjdjd"

	x := len(str)
	y := len(text)

	costMatrix := createRecurrentRelationMatrix(x, y)

	printRecurrentRelationMatrix(costMatrix)

	for i := 0; i <= y; i++ {
		rowInt(i, &costMatrix)
	}

	for i := 0; i <= x; i++ {
		columnInt(i, &costMatrix)
	}

	printRecurrentRelationMatrix(costMatrix)

	opt := make([]int, 3, 3)

	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {

			textChar := string(text[j-1])
			strChar := string(str[i-1])

			opt[MATCH] = costMatrix[i-1][j-1].Cost + match(strChar, textChar)
			opt[INSERT] = costMatrix[i][j-1].Cost + insert()
			opt[DELETE] = costMatrix[i-1][j].Cost + del()

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
	i, j := goalCell(costMatrix, str, text)

	fmt.Println("<i: ", i,"> <j: ",j,">")
	fmt.Println("Edit Cost = ", costMatrix[i][j].Cost)

	reconstructPath(costMatrix, str, text, i, j)
}