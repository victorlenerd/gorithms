package main

import (
	"fmt"
)

// SAMPLE 1 :- input := [][]int{{0,1}, {0,2}, {0,3}, {1,4}}
// SAMPLE 2 :- input := [][]int{{0,1}, {1,2}, {2,3}, {1,3}, {1,4}}
// SAMPLE 3 :- input := [][]int{{0,1}, {2,3}}

//4
//[[0,1],[2,3]]

// 2
// [[1,0]]

//3
//[[2,0],[2,1]]

func main() {
	input := [][]int{{1, 0}}
	n := 2
	graph := ConstructGraph(n, input)
	numberOfCycles, discoveredVertices := BFS(graph, 0)
	fmt.Println(graph)

	fmt.Println("discoveredVertices", discoveredVertices)
	fmt.Println("numberOfCycles", numberOfCycles)

	if discoveredVertices < n || numberOfCycles > 0 {
		fmt.Println("Not A Tree")
		return
	}

	fmt.Println("A Tree")
}

func ConstructGraph(n int, input [][]int) [][]int {
	adjList := make([][]int, n)

	for j := 0; j < len(input); j++ {
		ai := input[j][0]
		bi := input[j][1]

		if adjList[ai] != nil {
			adjList[ai] = append(adjList[ai], bi)
		} else {
			adjList[ai] = []int{bi}
		}

		if adjList[bi] != nil {
			adjList[bi] = append(adjList[bi], ai)
		} else {
			adjList[bi] = []int{ai}
		}
	}

	return adjList
}

func BFS(graph [][]int, start int) (int, int) {
	queue := []int{start}
	processed := []int{}

	cycles := 0

	for len(queue) > 0 {
		currentNode := queue[0]
		adjNodes := []int{}

		if Contains(processed, currentNode) {
			cycles++
		}

		// Are we at the end of the graph
		if len(graph) > currentNode {
			adjNodes = graph[currentNode]
		}

		for _, node := range adjNodes {
			if !Contains(processed, node) {
				queue = append(queue, node)
			}
		}

		processed = append(processed, currentNode)
		queue = queue[1:]
	}

	return cycles, len(processed)
}

func Contains(arr []int, n int) bool {
	for _, item := range arr {
		if item == n {
			return true
		}
	}
	return false
}