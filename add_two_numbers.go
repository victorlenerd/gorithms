package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) print() []int {
	r := []int{}

	var current *ListNode = l

	for current.Next != nil {
		r = append(r, current.Val)
		current = current.Next
	}

	r = append(r, current.Val)

	return r
}

func recurseNode(node *ListNode, stack *StackInt) {
	stack.push(node.Val)
	if node.Next != nil {
		recurseNode(node.Next, stack)
	} else {
		return
	}
}

func addNode(list *ListNode, val int) {
	var current *ListNode = list

	for current.Next != nil {
		current = current.Next
	}

	current.Next = &ListNode{Val: val}
}

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	var carry int = 0
//	var result ListNode = ListNode{}
//	curr, q, p := l1, l2, result
//
//	for p.Next != nil || q.Next != nil {
//		var x int = 0
//		var y int = 0
//
//		if p.Next != nil {
//			x = p.Val
//		}
//
//		if q.Next != nil {
//			y = q.Val
//		}
//
//		sum := carry + x + y
//		carry = sum / 10
//
//		curr.Next = &ListNode{ Val: sum % 10 }
//		curr = curr.Next
//
//		if p.Next != nil {
//			p = p.Next
//		}
//
//		if q.Next != nil {
//			q = q.Next
//		}
//
//	}
//
//	return result
//}
