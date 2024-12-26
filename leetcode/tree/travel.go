package main

import (
	"fmt"
)
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

//stack operations
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
func top(stack *[]*TreeNode) *TreeNode {
    return (*stack)[len(*stack)-1]
}

//recursive Traversal
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

//iterative Traversal
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

func inorderTraversal_iter(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    stack := []*TreeNode{}
    res := []int{}
    for len(stack) > 0 || root != nil {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        node := pop(&stack)
        res = append(res, node.Val)
        root = node.Right
    }
    return res
}

func postorderTraversal_iter(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    stack := []*TreeNode{root}
    res := []int{}
    mark := &TreeNode{}
    for len(stack) > 0 {
        for top(&stack).Left != nil && top(&stack).Left != mark && top(&stack).Right!= mark {
            push(&stack, top(&stack).Left)
        }
        if top(&stack).Right == nil || top(&stack).Right == mark {
            mark = top(&stack)
            res = append(res, pop(&stack).Val)
        } else {
            push(&stack, top(&stack).Right)
        }
    }
    return res 
}

func levelOrder(root *TreeNode) [][]int {
    queue := []*TreeNode{}
    res := [][]int{}
    if root == nil { return res }
    queue = append(queue, root)
    for len(queue) > 0 {
        level := []int{}
        for _, v := range queue {
            level = append(level, v.Val) 
            if v.Left != nil { queue = append(queue, v.Left) }
            if v.Right != nil { queue = append(queue, v.Right) }
            queue = queue[1:]
        }
        res = append(res, level)

    }
    return res
}




func printTree(root *TreeNode, level int) {
    if root == nil {
        return
    }
    // 先递归打印右子树，增加层级
    printTree(root.Right, level+1)
    // 打印空格体现层级关系
    for i := 0; i < level; i++ {
        fmt.Print("    ")
    }
    fmt.Printf("%d\n", root.Val)
    // 再递归打印左子树，增加层级
    printTree(root.Left, level+1)
}



func main() {
    root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, 
                              Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}}

    printTree(root, 0)
    pre_res1 := preorderTraversal_iter(root)
    pre_res2 := preorderTraversal(root)
    fmt.Println("preorder:")
    fmt.Println(pre_res1,"\n",pre_res2)

    in_res1 := inorderTraversal_iter(root)
    in_res2 := inorderTraversal(root)
    fmt.Println("inorder:")
    fmt.Println(in_res1,"\n",in_res2)

    post_res1 := postorderTraversal_iter(root)
    post_res2 := postorderTraversal(root)
    fmt.Println("postorder:")
    fmt.Println(post_res1,"\n",post_res2)

    level_res := levelOrder(root)
    fmt.Println("levelorder:")
    fmt.Println(level_res)    

}