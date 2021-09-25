package main

import (
	"fmt"
	"math"
)

func IsSolution(partialSolution []string, sizeOfCompleteSolution int) bool {
	return len(partialSolution) == sizeOfCompleteSolution
}

func ProcessSolution(subsets []string) {
	fmt.Print("{")
	for _, s := range subsets {
		fmt.Printf(" %v", s)
	}
	fmt.Print(" }\n")
}

func Contains(strs []string, s string) bool {
	for _, str := range strs {
		if str == s {
			return true
		}
	}

	return false
}

func ConstructCandidates(subset []string, sizeOfCompleteSet int) []string {
	numberOfCandidates := sizeOfCompleteSet - len(subset)
	candidates := []string{}

	potentialCandidates := []string{"A", "B", "C", "_"}

	for len(candidates) < numberOfCandidates {
		for _, pc := range potentialCandidates {
			if !Contains(subset, pc) && !Contains(candidates, pc) {
				candidates = append(candidates, pc)
			}
		}
	}

	return candidates
}

func Backtrack(subsets []string, sizeOfCompleteSet int) {
	if IsSolution(subsets, sizeOfCompleteSet) {
		ProcessSolution(subsets)
	} else {
		candidates := ConstructCandidates(subsets, sizeOfCompleteSet)
		for i := 0; i < len(candidates); i++ {
			// Extend partial solution with candidates
			partialSolution := append(subsets, candidates[i])
			Backtrack(partialSolution, sizeOfCompleteSet)
		}
	}
}

func generateSubsets() {
	str := []string{"A", "B", "C"}

	sizeOfSet := len(str)
	maxSubsets := int(math.Pow(2, float64(sizeOfSet)))
	subsets := make([]string, 0, maxSubsets)

	Backtrack(subsets, sizeOfSet)

	fmt.Println(subsets)
}


func main() {
	generateSubsets()
}
