package main

import (
	"fmt"
	"strconv"
	"strings"
)

const Alphas string = "abcdefghijklmnopqrstuvwxyz"

func GenerateDigitAlphas(str string) string {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	if num == 1 {
		return  ""
	}

	if num == 7 {
		startIndex := 15
		endIndex := startIndex + 4
		return Alphas[startIndex: endIndex]
	}

	if num == 8 {
		startIndex := 19
		endIndex := startIndex + 3
		return Alphas[startIndex: endIndex]
	}

	if num == 9 {
		startIndex := 22
		endIndex := startIndex + 4
		return Alphas[startIndex: endIndex]
	}

	startIndex :=  (3 * num) - 6
	endIndex := startIndex + 3

	return Alphas[startIndex: endIndex]
}

func GetSizeOfCompleteSet(inputStr string) int {
	seen := []string{}
	completeSet := 1

	for index, _ := range inputStr {
		input := string(inputStr[index])
		if !strings.Contains(strings.Join(seen, ""), input) {
			numChars := GenerateDigitAlphas(input)
			completeSet = completeSet * len(numChars)
		}
	}

	return completeSet
}

func IsCompleteSubset(subset []string, sizeOfCompleteSet int) bool {
	return len(subset) == sizeOfCompleteSet
}

func ConstructCandidates(inputStr string, partialSet []string) []string {
	candidates := []string{}
	currentPosition := len(partialSet)

	nextDigit := inputStr[currentPosition]
	numChars := GenerateDigitAlphas(string(nextDigit))

	for i, _ := range numChars {
		s := string(numChars[i])
		candidates = append(candidates, s)
	}

	return candidates
}

func AddSubsetToCompleteSet(partialSet []string, completeSets *[]string)  {
	*completeSets = append(*completeSets, strings.Join(partialSet, ""))
}

func Backtrack(inputStr string, partialSet []string, completeSets *[]string, sizeOfSubset int, sizeOfCompleteSet int) {
	if len(*completeSets) == sizeOfCompleteSet {
		return
	}

	if IsCompleteSubset(partialSet, sizeOfSubset) {
		AddSubsetToCompleteSet(partialSet, completeSets)
	} else {
		candidates := ConstructCandidates(inputStr, partialSet)
		for _, candidate := range candidates {
			extendedSet := append(partialSet, candidate)
			Backtrack(inputStr, extendedSet, completeSets, sizeOfSubset, sizeOfCompleteSet)
		}
	}
}

func main() {
	inputStr := "22"
	characters := []string{}

	for index, _ := range inputStr {
		numChars := GenerateDigitAlphas(string(inputStr[index]))

		for i, _ := range numChars {
			s := string(numChars[i])
			if !strings.Contains(strings.Join(characters, ""), s) {
				characters = append(characters, s)
			}
		}
	}

	sizeOfSubset := len(inputStr)
	sizeOfCompleteSet := GetSizeOfCompleteSet(inputStr)

	completeSet := []string{}
	partialSet := []string{}

	Backtrack(inputStr, partialSet, &completeSet, sizeOfSubset, sizeOfCompleteSet)

	fmt.Println(GenerateDigitAlphas("8"))
	fmt.Println(completeSet)
}
