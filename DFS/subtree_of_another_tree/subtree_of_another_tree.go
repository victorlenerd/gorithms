// Leetcode: https://leetcode.com/problems/subtree-of-another-tree/

package main

import "fmt"

//Input #1
var root = []interface{}{4,5}
var subRoot = []interface{}{4, nil, 5}

type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}

func AddChildren(node *TreeNode, input []interface{}, index int) {
	leftIndex := index + 1
	rightIndex := index + 2

	if leftIndex < len(input) && input[leftIndex] != nil  {
		node.Left = &TreeNode{
			Val: input[leftIndex].(int),
		}

		AddChildren(node.Left, input, leftIndex * 2)
	}

	if rightIndex < len(input) && input[rightIndex] != nil  {
		node.Right = &TreeNode{
			Val: input[rightIndex].(int),
		}

		AddChildren(node.Right, input, rightIndex * 2)
	}
}

func ConstructBinaryTree(input []interface{}) *TreeNode {
	node := TreeNode{
		Val: input[0].(int),
	}

	AddChildren(&node, input, 0)

	return  &node
}

func ContainsSubTree(root *TreeNode, subRoot *TreeNode) bool {
	if root.Val == subRoot.Val {
		leftValidity := true
		rightValidity := true

		if root.Left != nil && subRoot.Left != nil {
			leftValidity =  ContainsSubTree(root.Left, subRoot.Left)
		}

		if root.Right != nil && subRoot.Right != nil {
			rightValidity = ContainsSubTree(root.Right, subRoot.Right)
		}

		if root.Right != nil && subRoot.Right == nil {
			rightValidity = false
		}

		if root.Right == nil && subRoot.Right != nil {
			rightValidity = false
		}

		if root.Left != nil && subRoot.Left == nil {
			leftValidity = false
		}

		if root.Left == nil && subRoot.Left != nil {
			leftValidity = false
		}

		return rightValidity && leftValidity
	}


	return false
}

func main()  {
	rootNode := ConstructBinaryTree(root)
	subRootNode := ConstructBinaryTree(subRoot)

	rootStack := []TreeNode{*rootNode}

	for len(rootStack) > 0 {
		currentNode := rootStack[0]
		rootStack = rootStack[1:]

		if currentNode.Left != nil {
			rootStack = append([]TreeNode{*currentNode.Left}, rootStack...)
		}

		if currentNode.Right != nil {
			rootStack = append([]TreeNode{*currentNode.Right}, rootStack...)
		}

		if currentNode.Val == subRootNode.Val {
			if ContainsSubTree(&currentNode, subRootNode) {
				fmt.Println("Contain Subtree")
				return
			}
		}
	}

	fmt.Println("Does not contain subtree")
}

