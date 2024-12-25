package main

import (
	"fmt"
)


func Swap(a, b *int) {*a, *b = *b, *a}


// func maxHeapify(arr []int, i, length int){
// 	dad := i
// 	lson := i * 2 + 1
// 	rson := i * 2 + 2
// 	if rson < length && arr[dad] < arr[lson] {
// 		dad = lson
// 	}
// 	if rson < length && arr[dad] < arr[rson] {
// 		dad = rson
// 	}
// 	if dad != i {
// 		Swap(&arr[dad], &arr[i])
// 		maxHeapify(arr, dad, length)
// 	}
// }

func minHeapify(arr []int, i, length int) {
	dad := i
	lson := i * 2 + 1
	rson := i * 2 + 2
	if rson < length && arr[dad] > arr[lson] {
		dad  = lson
	}
	if rson < length && arr[dad] > arr[rson] {
		dad = rson
	}
	if dad != i {
		Swap(&arr[dad], &arr[i])
		minHeapify(arr, dad, length)
	}
}


func heapsort(arr []int, length int) {
	for i := length/2-1; i >= 0; i-- {
		minHeapify(arr, i, length)
	}
	for i := length-1; i >= 0; i-- {
		Swap(&arr[i], &arr[0])
		minHeapify(arr, 0, i)
	}
}



func main()  {
	arr := []int{20, 40, 30, 88, 14, 56, 24, 77, 65}
	length := len(arr)
	heapsort(arr, length)
	fmt.Println(arr)
}