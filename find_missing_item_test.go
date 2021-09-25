package main

import (
	"fmt"
	"testing"
)

func TestFindMissingItem(t *testing.T) {

	t.Log("Find missing item test")
	{
		a := []int{1, 2, 3, 4, 6}
		b := []int{2, 3, 4}

		c := FindMissingItem(a, b)

		fmt.Print(c)
	}
}
