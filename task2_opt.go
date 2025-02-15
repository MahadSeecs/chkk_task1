package main

import (
	"fmt"
	"sort"
)

//func to reverse sorted second half after mid
func reverse(arr []int) {
	for left, right := 0, len(arr)-1; left < right; left, right = left+1, right-1 {
		arr[left], arr[right] = arr[right], arr[left]
	}
}


func Wave(arr []int, x int) []int {
	size := len(arr)

	for i := 0; i < size; i += (2*x + 1) {
		start := i
		end := i + (2*x + 1)
		if end > size {
			end = size
		}

		mid := start + (end-start)/2

		//sort the extracted block
		sort.Ints(arr[start:end])

		//swap max element to middle
		arr[mid], arr[end-1] = arr[end-1], arr[mid]

		//reverse the second half in descending order
		reverse(arr[mid+1:end])
	}

	return arr
}

// func main() {
// 	arr := []int{3,6,5,10,7,11,1,2,9,4}
// 	x := 2

// 	result := Wave(arr, x)
// 	fmt.Println(result)
// }
