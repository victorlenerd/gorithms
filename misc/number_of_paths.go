package misc

import "fmt"

type Grid struct {
	size int
	grid [][]int
}

const (
	EAST  = 0
	NORTH = 1
)

func (g *Grid) make(size int) [][]int {
	grid := make([][]int, size)

	for i := range grid {
		innerGrid := make([]int, size)
		for j := range grid {
			innerGrid[j] = i + j
		}
		grid[i] = innerGrid
	}

	g.grid = grid

	return g.grid
}

func (g *Grid) print() {
	for _, g := range g.grid {
		fmt.Println(g)
	}
}

// Validates step towards a direction
func canMove(n int, x int, y int, ls Direction, ns Direction) bool {
	if x >= y && (ls != E && ns != E) {
		return false
	}

	return true
}

type Direction int

const (
	N Direction = 0
	E Direction = 1
)

type Step struct {
	x         int
	y         int
	direction Direction
}

func step(destination int, x int, y int, paths int, dir Direction) int {
	if x == destination && y == destination {
		return paths + 1
	}

	if canMove(destination, x+1, y, dir, N) {
		return step(destination, x+1, y, paths, N)
	}

	if canMove(destination, x, y+1, dir, E) {
		return step(destination, x, y+1, paths, E)
	}

	return paths
}

func FindNumberOfPaths(p int) int {
	var destination int = p - 1
	//var seen map[string]int

	return step(destination, 0, 0, 0, E)
}
