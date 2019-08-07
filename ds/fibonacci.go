package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	defer timeTrack(time.Now(), "Main")
	f := fibonacci()
	for i := 0; i < 100; i++ {
		fmt.Println(f())
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

/*
Solution 1:
func fibonacci() func() uint {
	var v1, v2 uint = 0, 1
	return func() uint {
		v2 = v2 + v1
		v1 = v2 - v1
		return v2 - v1
	}
}
*/

/*
Solution 2:
func fibonacci() func() int {
	v1, v2 := 0, 1

	return func() int {
		defer func() {
			v1, v2 = v2, v1+v2
		}()

		return v1
	}
}*/
