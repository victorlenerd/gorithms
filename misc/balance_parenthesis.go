package misc

func balancedParenthesis(parenthesis []string) bool {
	stack := Stack{}

	for _, parent := range parenthesis {
		if parent == "(" {
			stack.push(parent)
		} else {
			stack.pop()
		}
	}

	return stack.len() == 0
}

func matchedBalancedParenthesis(parenthesis []string) bool {
	stack := Stack{}

	for _, parent := range parenthesis {
		if parent == "(" || parent == "[" || parent == "{" {
			stack.push(parent)
		} else {
			match := stack.top()

			if match == "(" && parent == ")" {
				stack.pop()
			}

			if match == "{" && parent == "}" {
				stack.pop()
			}

			if match == "[" && parent == "]" {
				stack.pop()
			}
		}
	}

	return stack.len() == 0
}
