
# Tree 
## 三种遍历方法
- 简介写法
```go
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
```
> 注意在append中，常规写法是`append([]int, int)`，但是如果第二个参数是切片的话，需要加上`...`，否则会报错，例如`append([]int, []int...)`
> 

非递归的版本

```go
//前序遍历
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
func preorderTraversal(root *TreeNode) []int {
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
```
>可以作为stack实现的魔板
`*[]TreeNode` 和 `[]*TreeNode`
*[]TreeNode表示指向切片的一个指针，
[]*TreeNode表示一个切片，里面的元素是指向TreeNode的指针
当然在二叉树的这种形态下，使用第二种会更好，当需要存储树节点的引用，并在多个地方共享这些节点时，使用 []*TreeNode。
| **特性**             | `*[]TreeNode`                     | `[]*TreeNode`                     |
|----------------------|------------------------------------|------------------------------------|
| **数据类型**         | 指向切片的指针                    | 指针切片                          |
| **切片的内容**       | 存储 `TreeNode` 的值              | 存储指向 `TreeNode` 的指针        |
| **内存占用**         | 每个元素占用 `TreeNode` 的内存     | 每个元素仅占用指针大小（8 字节）  |
| **是否共享数据**     | 对切片的内容修改不会共享           | 对指针指向的对象修改会共享        |
| **适用场景**         | 修改切片本身（例如添加/删除元素）  | 操作共享数据（如树节点）          |


