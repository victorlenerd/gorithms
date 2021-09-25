/*
	Title: Generate Parenthesis
	Author: Victor Nwaokocha
*/

package main

func generateParenthesis(n int) []string {
	results := []string{}
	generateBalancedParenthesis(n, n, "", &results)
	return results
}

func generateBalancedParenthesis(
	numOfLeftPNeeded int,
	numOfRightPNeeded int,
	parentsProgress string,
	results *[]string) {

	if numOfLeftPNeeded == 0 && numOfRightPNeeded == 0 {
		*results = append(*results, parentsProgress)
		return
	}

	if numOfLeftPNeeded > 0 {
		generateBalancedParenthesis(
			numOfLeftPNeeded-1,
			numOfRightPNeeded,
			parentsProgress+"(",
			results)
	}

	if numOfLeftPNeeded < numOfRightPNeeded {
		generateBalancedParenthesis(
			numOfLeftPNeeded,
			numOfRightPNeeded-1,
			parentsProgress+")",
			results)
	}
}
