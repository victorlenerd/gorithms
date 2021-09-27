// Leetcode: https://leetcode.com/problems/subtree-of-another-tree/

package main

import "fmt"

//Input #1
var root = []interface{}{3,4,5,1,2,nil,nil,nil,nil,0}
var subRoot = []interface{}{4,1,2}

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

func main()  {


	node := ConstructBinaryTree(root)
	fmt.Println("Node: ", node)

}