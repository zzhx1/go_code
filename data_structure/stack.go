package main

import (
	"fmt"
)


func NewStack() []int { return []int{} } 
func push() {}

func main() {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println("Stack size:", stack.Size())
	fmt.Println("Peek:", stack.Peek())
	fmt.Println("Pop:", stack.Pop())
	fmt.Println("Stack size after pop:", stack.Size())
}


func initStack() []int {
	return []int{}
}


func main()  {
	stack := initStack()
	
}