package main

import "fmt"
import "strconv"



func calculate(s string) int {
	// fmt.Println(s)
	res, _ := recur(s, 0)
	return res
}
func get_op(op string, num int) int {
	if op == "-" {
		return -num
	}
	return num
}
func recur(s string, index int) (int, int) { //两个返回值一个是res， 一个右括号的索引
    res := 0
	fmt.Println(s)
    for i := range(len(s)) {
		
        if s[i] == '+' || s[i] == '-' || s[i] == ' '{
            continue
        }
        if s[i] == ')' { //遇到右括号递归返回
            return res, i+1
        }
        if s[i] == '(' { //遇到左括号向下递归
            sub_res, rightIdx := recur(s[i+1:], i)
            if i == 0{
                res += sub_res
            }else{
                res += get_op(string(s[i-1]), sub_res)
            }
            // res += get_op(s[i-1], sub_res)
            i += rightIdx
        }
        if s[i] == '+' || s[i] == '-' || s[i] == ' '{
            continue
        }
        numstr := string(s[i])
		op_index := i-1
        i++
        for i < len(s) && s[i] >= '0' && s[i] <= '9' {
            numstr += string(s[i])
            i++
        }
        num, _ := strconv.Atoi(numstr)
		
        res += get_op(string(s[op_index]), num)
        i--
    }
    return res, 0
}

func main()  {
	string1 := "1 + 1"			//2
	string2 := " 2-1 + 2 "		//3
	string3 := "(1+(4+5+2)-3)+(6+8)"//23
	string4 := "2147483647"//

	fmt.Println(calculate(string1))
	fmt.Println(calculate(string2))
	fmt.Println(calculate(string3))
	fmt.Println(calculate(string4))

}