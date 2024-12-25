package main
import "fmt"




//go语言在切片的最前方添加元素
// 例如将 a 添加到 arr 的最前方
// _a := []int{a}
// arr = append(_a, arr...)
// 必须加上...
/*
注意逻辑：
1. stack不需要单独的存储所有的数字，存储所有的索引就可以
2. 中间出现一个温度到最后都是比他低的，那就必须返回0

{89,62,70,58,47,47,46,76,100,70}
[8,1,5,4,3,2,1,1,0,0]
*/
// func dailyTemperatures(temperatures []int) []int {
//     len_temp := len(temperatures)
// 	stack := []int{} //存储temperature的索引
// 	res := make([]int, len_temp)

// 	stack = append(stack, len(temperatures)-1)
// 	res[len(res)-1] = 0
// 	for i := len_temp-2; i >= 0; i-- {
// 		day := 0
// 		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack) - 1]] {
// 			day = stack[len(stack)-1] - i
// 			stack = stack[:len(stack) - 1]
// 			// res[len()]
// 		}
// 		if len(stack) == 0 {
// 			day = 0
// 		}else {
// 			day++
// 		}
// 		stack = append(stack, i)
// 		res[i] = day
// 	}
// 	return res
// }
func dailyTemperatures(temperatures []int) []int {
    length := len(temperatures)
	stack := []int{} //存储索引
	res := make([]int, length) //输出结果
	// 2 4 1 3 [1, 0, 1, 0] 
	for i := 0; i < length; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return res
}



// 正确代码
// func dailyTemperatures(temperatures []int) []int {
// 	length := len(temperatures)
// 	res := make([]int, length)
// 	stack := []int{} // 栈存储索引

// 	for i := 0; i < length; i++ {
// 		// 维护单调递减栈
// 		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
// 			top := stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]
// 			res[top] = i - top // 当前索引减去栈顶索引
// 		}
// 		stack = append(stack, i)
// 	}

// 	// 栈中剩余元素对应结果为 0（无需操作，默认值已是 0）
// 	return res
// }


func main() {

	temperatures := []int{89,62,70,58,47,47,46,76,100,70}
	// [1,1,4,2,1,1,0,0]
	res := dailyTemperatures(temperatures)
	fmt.Println(res)

}



