// Linear Search in Golang
package main

import "fmt"

func linearSearch(arr []int, search int) bool {
	for _, v := range arr {
		if v == search {
			return true
		}
	}
	return false
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 23, 67, 88, 32, 5, 87, 90, 12}
	if true == linearSearch(arr, 123) {
		fmt.Println("Record Found")
	} else {
		fmt.Println("Record not Found")
	}
}
