package main

import (
	"fmt"
)
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
     }
func preorderTraversal(root *TreeNode) []int {
    if root == nil { return nil }
    return append(append([]int{root.Val}, preorderTraversal(root.Left)...), preorderTraversal(root.Right)...)
}
func inorderTraversal(root *TreeNode) []int {
    if root == nil { return nil }  
    return append(append(inorderTraversal(root.Left), root.Val), inorderTraversal(root.Right)...)
}
func postorderTraversal(root *TreeNode) []int {
    if root == nil { return nil }
    return append(append(postorderTraversal(root.Left), postorderTraversal(root.Right)...), root.Val)
}


 func pop(stack *[]*TreeNode) *TreeNode {
    node := (*stack)[len(*stack)-1]
    *stack = (*stack)[:len(*stack)-1]
    return node
}

func push(stack *[]*TreeNode, node *TreeNode) {
    if node != nil { 
        *stack = append(*stack, node)
    }
}

func preorderTraversal_iter(root *TreeNode) []int {
    if root == nil { 
        return []int{}
    }
    stack := []*TreeNode{root} 
    res := []int{}
    for len(stack) > 0 {
        node := pop(&stack)
        res = append(res, node.Val)
        push(&stack, node.Right) 
        push(&stack, node.Left)  
    }
    return res
}

