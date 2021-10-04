package misc

import (
	"fmt"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {

	t.Log("Generate N numbers of parenthesis")
	{
		parenthesis := generateParenthesis(2)
		fmt.Println(parenthesis)
	}
}
