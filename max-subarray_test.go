package main

import "testing"

func TestMaxSubArray(t *testing.T) {

	t.Log("Max Subarray")
	{
		arr := []float64{-2.0, -3.0, 4.0, -1.0, -2.0, 1, 5, -3}
		max := maxSubarray(arr)

		if max != 15 {
			t.Fatalf("\t Max sub array should be %v but got %v", 15, max)
		}
	}

}
