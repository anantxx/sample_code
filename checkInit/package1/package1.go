package package1

import "fmt"

var v1 = displayS()

func displayS() int {
	fmt.Println("Package 1 Display")
	return 11
}

func init() {
	fmt.Println("Package 1 Init : ", v1)
}

func DisplayAfter() {
	fmt.Println("Package 1 Display After", v1)
}
