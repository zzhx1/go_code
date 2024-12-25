
type KeyValue struct {
	key   int
	value int
}

func Swap(a, b *KeyValue) {
	*a, *b = *b, *a
}

func minHeapify(arr *[]KeyValue, i, length int) {
	dad := i
	lson := 2*i + 1
	rson := 2*i + 2
	if lson < length && (*arr)[dad].value > (*arr)[lson].value {
		dad = lson
	}
	if rson < length && (*arr)[dad].value > (*arr)[rson].value {
		dad = rson
	}
	if dad != i {
		Swap(&(*arr)[dad], &(*arr)[i])
		minHeapify(arr, dad, length)
	}
}

func pop(priority_queue []KeyValue) []KeyValue {
	return priority_queue[:len(priority_queue)-1]
}

func push(priority_queue []KeyValue, a KeyValue) []KeyValue {
	return append(priority_queue, a)
}

func topKFrequent(nums []int, k int) []int {
	hashmap := make(map[int]int)
	for _, v := range nums {
		hashmap[v]++
	}
	priority_queue := make([]KeyValue, 0, k+1)
	for key, value := range hashmap {
		priority_queue = push(priority_queue, KeyValue{key, value})
		for i := len(priority_queue)/2-1; i >= 0; i-- {
			minHeapify(&priority_queue, i, len(priority_queue))
		}
		if len(priority_queue) > k {
			Swap(&priority_queue[0], &priority_queue[len(priority_queue)-1])
			priority_queue = pop(priority_queue)
			minHeapify(&priority_queue, 0, len(priority_queue))
		}
	}
	res := make([]int, 0, k)
	for i := k - 1; i >= 0; i-- {
		res = append(res, priority_queue[i].key)
	}

	return res
}










// package main

// import (
//     "fmt"
//     // "sort"
// )

// // KeyValue结构体用于存储map中的键值对
// type KeyValue struct {
//     Key   string
//     Value int
// }

// // ByValue结构体切片类型，用于实现sort.Interface接口以按值排序
// type ByValue []KeyValue

// func (a ByValue) Len() int { return len(a) }

// // Less方法定义按照值比较的规则，用于排序判断
// func (a ByValue) Less(i, j int) bool { return a[i].Value < a[j].Value }

// func (a ByValue) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// func main() {
//     m := map[string]int{"c": 5, "a": 3, "b": 1, "d": 10, "e": 11}
//     kv := make(ByValue, 0, len(m))
//     // 将map中的键值对存入kv切片
//     for k, v := range m {
//         kv = append(kv, KeyValue{k, v})
//     }
//     sort.Sort(ByValue(kv))
//     // 遍历排序后的切片输出相应的键值对
//     for _, kvPair := range kv {
//         fmt.Printf("Key: %s, Value: %d\n", kvPair.Key, kvPair.Value)
//     }
// }





package main

import "fmt"

type KeyValue struct {
	key   int
	value int
}

func Swap(a, b *KeyValue) {
	*a, *b = *b, *a
}


func minHeapify(arr *[]KeyValue, i, length int) {
	dad := i
	lson := 2*i + 1
	rson := 2*i + 2
	if lson < length && (*arr)[dad].value > (*arr)[lson].value {
		dad = lson
	}
	if rson < length && (*arr)[dad].value > (*arr)[rson].value {
		dad = rson
	}
	if dad != i {
		Swap(&(*arr)[dad], &(*arr)[i])
		minHeapify(arr, dad, length)
	}
}

func pop(priority_queue []KeyValue) []KeyValue {
	return priority_queue[:len(priority_queue)-1]
}

func push(priority_queue []KeyValue, a KeyValue) []KeyValue {
	return append(priority_queue, a)
}

func topKFrequent(nums []int, k int) []int {
	hashmap := make(map[int]int)
	for _, v := range nums {
		hashmap[v]++
	}
	priority_queue := make([]KeyValue, 0, k+1)
	for key, value := range hashmap {
		priority_queue = push(priority_queue, KeyValue{key, value})
		for i := len(priority_queue)/2-1; i >= 0; i-- {
			minHeapify(&priority_queue, i, len(priority_queue))
		}
		if len(priority_queue) > k {
			Swap(&priority_queue[0], &priority_queue[len(priority_queue)-1])
			priority_queue = pop(priority_queue)
			minHeapify(&priority_queue, 0, len(priority_queue))
		}
	}
	res := make([]int, 0, k)
	for i := k - 1; i >= 0; i-- {
		res = append(res, priority_queue[i].key)
	}

	return res
}

func main() {
	a := []int{5,3,1,1,1,3,5,73,1}

	b := topKFrequent(a, 3)
	fmt.Println(b) 
}



package main

import (
	"fmt"
)




// Stack 定义一个栈结构体
type Stack struct {
	items []int
}

// NewStack 创建一个新的空栈
func NewStack() *Stack {
	return &Stack{items: make([]int, 0)}
}

// Push 将元素压入栈顶
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop 从栈顶弹出元素
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic("Stack is empty")
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item
}

// Peek 查看栈顶元素，但不弹出
func (s *Stack) Peek() int {
	if s.IsEmpty() {
		panic("Stack is empty")
	}
	return s.items[len(s.items)-1]
}

// IsEmpty 检查栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size 获取栈中元素的数量
func (s *Stack) Size() int {
	return len(s.items)
}

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