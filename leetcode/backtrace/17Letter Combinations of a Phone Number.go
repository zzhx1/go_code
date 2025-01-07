
package main
import (
	"fmt"
)



func letterCombinations(digits string) (res []string) {
	mp := map[byte]string{
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }
	 // temp := []byte //注意字符数组的初始化过程
	var temp []byte
	var dfs func(int )
	dfs = func(idx int) {
		if idx == len(digits) {
			res = append(res, string(temp))//字符数组转变为string
			return
		}
		letter := mp[(digits[idx])]
		
		for _, str := range letter { //巨大发现，letter是string类型的，但是如果我用 for _, val range的话竟然val是int类型的ASCIIma
			fmt.Println(str)
			temp = append(temp, byte(str)) // 
			dfs(idx + 1)
			temp = temp[:len(temp)-1] // 回溯
		}  
	}
	dfs(0)
	return res
}
//在Go语言中，没有像Python中的type()函数那样直接获取变量类型的内置函数。
//但是，你可以使用fmt.Printf函数结合%T格式化动词来打印变量的类型。

func main() {
    digits := "23"
    result := letterCombinations(digits)
    fmt.Println(result) // 输出 ["ad" "ae" "af" "bd" "be" "bf" "cd" "ce" "cf"]
}
