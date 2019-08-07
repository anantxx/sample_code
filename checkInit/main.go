package main

import (
	"fmt"

	"github.com/anant/final/checkInit/package1"
	"github.com/anant/final/checkInit/package2"
)

var v1 = displayS()

func displayS() int {
	fmt.Println("Main Display")
	return 33
}

func init() {
	fmt.Println("Main Init : ", v1)
}

func main() {
	package1.DisplayAfter()
	package2.DisplayAfter()
}

func init() {
	fmt.Println("Main Init 2: ", v1)
}
