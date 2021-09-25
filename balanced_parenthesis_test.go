package main

import (
	"strings"
	"testing"
)

func TestBalancedParenthesis(t *testing.T) {

	t.Log("Balanced Parenthesis")
	{

		t.Logf("\t Unbalanced Parenthesis")
		{
			badParenthesis := "((())"
			isUnbalanced := balancedParenthesis(strings.Split(badParenthesis, ""))

			if isUnbalanced {
				t.Logf("\t Expected %v but got %v", false, isUnbalanced)
			}
		}

		t.Logf("\t Balanced Parenthesis")
		{
			goodParenthesis := "(())()"
			isBalanced := balancedParenthesis(strings.Split(goodParenthesis, ""))

			if !isBalanced {
				t.Logf("\t Expected %v but got %v", true, isBalanced)
			}
		}
	}

	t.Log("Matched Balanced Parenthesis")
	{
		t.Logf("\t Matched Unbalanced Parenthesis")
		{
			badParenthesis := "((())}[]}{}"

			isUnbalanced := matchedBalancedParenthesis(strings.Split(badParenthesis, ""))

			if isUnbalanced {
				t.Logf("\t Expected %v but got %v", false, isUnbalanced)
			}
		}

		t.Logf("\t Matched Balanced Parenthesis")
		{
			goodParenthesis := "{(())(){}[]}"
			isBalanced := matchedBalancedParenthesis(strings.Split(goodParenthesis, ""))

			if !isBalanced {
				t.Logf("\t Expected %v but got %v", true, isBalanced)
			}
		}
	}
}
