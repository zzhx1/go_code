package main

import (
	"fmt"
)


func NewStack() []int { return []int{} } 
func push(stack *[]int, a int) { *stack = append(*stack, a) }
func pop(stack *[]int) int {
	length := len(*stack)
	res := (*stack)[length-1]
	*stack = (*stack)[:length-1]
	return res
}




func main()  {
	stack := initStack()
	
}