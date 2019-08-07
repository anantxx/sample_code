package package2

import "fmt"

var v1 = displayS()

func displayS() int {
	fmt.Println("Package 2 Display")
	return 22
}

func init() {
	fmt.Println("Package 2 Init : ", v1)
}

func DisplayAfter() {
	fmt.Println("Package 1 Display After", v1)
}
