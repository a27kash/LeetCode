package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	canBeRoot := make(map[int]bool)
	dict := make(map[int]*TreeNode)
	for _, x := range descriptions {
		parent := x[0]
		child := x[1]
		isLeft := x[2]
		var parentNode *TreeNode
		var childNode *TreeNode
		if p, ok := dict[parent]; ok {
			parentNode = p
		} else {
			parentNode = &TreeNode{
				Val: parent,
			}
			dict[parent] = parentNode
			// add to can be root
			canBeRoot[parent] = true
			// fmt.Println(canBeRoot)
		}
		if c, ok := dict[child]; ok {
			childNode = c
			// remove from can be root
			delete(canBeRoot, child)
			// fmt.Println(canBeRoot)
		} else {
			childNode = &TreeNode{
				Val: child,
			}
			dict[child] = childNode
		}
		// link nodes
		if isLeft == 0 {
			parentNode.Right = childNode
		} else {
			parentNode.Left = childNode
		}
	}
	var root int
	for k := range canBeRoot {
		root = k
	}
	return dict[root]
}

func main() {
	fmt.Println("hello, world!")

	fmt.Println(createBinaryTree(
		[][]int{
			[]int {
				20,15,1,
			},
			[]int {
				20,17,0,
			},
			[]int {
				50,20,1,
			},
			[]int {
				50,80,0,
			},
			[]int {
				80,19,1,
			},
		},
	))
}
