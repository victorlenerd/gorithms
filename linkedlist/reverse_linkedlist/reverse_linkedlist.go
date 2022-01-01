package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {

	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	current := head.Next
	prev := head

	prev.Next = nil

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	fmt.Println(prev)
}