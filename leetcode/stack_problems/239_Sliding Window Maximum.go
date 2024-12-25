package main

import "fmt"


type max_num struct {
    num   int
    index int
}
func maxSlidingWindow(nums []int, k int) []int {
    res := []int{}
    stack := []int{}

    for i, v := range nums {
        
        for len(stack) > 0 && v > nums[stack[len(stack)-1]] {
            stack = stack[:len(stack)-1]
        }
        
        
        stack = append(stack , i)
        if  i - stack[0] > k {
            stack = stack[1:]
        }
        if i >= k-1 {
            res = append(res, nums[stack[0]])
        }

    }

    return res
}

func main() {
	nums := []int{1,3,1,2,0,5}
	k := 3
	a := maxSlidingWindow(nums, k)
    fmt.Println(a)
}//[3,3,2,5]