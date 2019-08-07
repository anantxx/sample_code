// Linear Search in Golang
package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, searchVal int) bool {
	sort.Ints(arr)
	low := 0
	high := len(arr) - 1
	for low <= high {
		median := (low + high) / 2
		if arr[median] < searchVal {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(arr) || arr[low] != searchVal {
		return false
	}

	return true
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 23, 67, 88, 32, 5, 87, 90, 12, 2}
	if true == binarySearch(arr, 2) {
		fmt.Println("Record Found")
	} else {
		fmt.Println("Record not Found")
	}
}
