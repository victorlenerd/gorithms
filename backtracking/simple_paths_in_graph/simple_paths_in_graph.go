package main

import (
	"fmt"
)

type Backtrack struct {
	AllPaths [][]int
}

func main()  {
	backtrack := Backtrack{
		AllPaths: [][]int{},
	}
	input := [][]int{{2}, {}, {1}}
	potentialCompletePath := []int{0}
	backtrack.Backtrack(potentialCompletePath, input)

	fmt.Println(backtrack.AllPaths)
}

func (backtrack *Backtrack) Backtrack(potentialCompletePath []int, input [][]int) {
	if backtrack.IsCompletePath(potentialCompletePath, input) {
		backtrack.ProcessCompletePath(potentialCompletePath)
	} else {
		nodes := backtrack.ConstructPotentialPath(potentialCompletePath, input)
		for _, node := range nodes {
			backtrack.Backtrack(append(potentialCompletePath, node), input)
		}
	}
}

func (backtrack *Backtrack) IsCompletePath(potentialCompletePath []int, input [][]int) bool {
	potentialCompletePathLen := len(potentialCompletePath)
	if potentialCompletePathLen < 1 {
		return false
	}

	lastNode := potentialCompletePath[potentialCompletePathLen - 1]

	if lastNode == len(input) - 1 {
		return true
	}

	return false
}

func (backtrack *Backtrack) ProcessCompletePath(potentialCompletePath []int)  {
	values := []int{}
	for _, p := range potentialCompletePath {
		values = append(values, p)
	}
	backtrack.AllPaths = append(backtrack.AllPaths, values)
}

func (backtrack *Backtrack) ConstructPotentialPath(potentialCompletePath []int, input [][]int) []int {
	potentialCompletePathLen := len(potentialCompletePath)

	if potentialCompletePathLen < 1 {
		return input[0]
	}

	lastNode := potentialCompletePath[potentialCompletePathLen - 1]
	adjNodes := input[lastNode]

	return adjNodes
}