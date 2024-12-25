package main

import (
	"fmt"
	// "sort"
)



//暴力解法
// type KeyValue struct {
//     Key   int
//     Value int
// }
// // ByValue结构体切片类型，用于实现sort.Interface接口以按值排序
// type ByValue []KeyValue

// func (a ByValue) Len() int { return len(a) }

// // Less方法定义按照值比较的规则，用于排序判断
// func (a ByValue) Less(i, j int) bool { return a[i].Value < a[j].Value }

// func (a ByValue) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// func topKFrequent(nums []int, k int) []int {
//     hashmap := make(map[int]int)
//     res := []int{}
//     for _, v := range nums {
//         hashmap[v]++
//     }
// 	fmt.Println(hashmap)

//     kv := make([]KeyValue,0, len(hashmap)) //make出来的数组，后面要append要注意设置初始长度为0
//     for k, v := range hashmap {
//         kv = append(kv, KeyValue{k, v})
//     }
// 	fmt.Println()
//     sort.Sort(ByValue(kv))

//     for i := range k {
//         res = append(res, kv[len(kv)-i-1].Key)
//     }

    
//     return res
// }

//堆写法

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



func main()  {
	a := []int{1,1,1,2,2,2,2,3,1,5,5,6,6,6,6,1,2,2,3,7,7,9,9,9,9,9,9,9}

	b := topKFrequent(a, 3)
	fmt.Println(b)


}