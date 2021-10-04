package misc

import "testing"

func TestStack(t *testing.T) {

	t.Log("Stack Implementation")
	{
		stack := Stack{}

		stack.push("A")
		stack.push("B")
		stack.push("C")

		c := stack.pop()
		if c != "C" {
			t.Fatalf("\t Expected poped value to be %v but got %v", "C", c)
		}

		b := stack.pop()
		if b != "B" {
			t.Fatalf("\t Expected poped value to be %v but got %v", "B", b)
		}

		a := stack.pop()
		if a != "A" {
			t.Fatalf("\t Expected poped value to be %v but got %v", "A", a)
		}
	}

}
