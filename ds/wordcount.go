package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	defer timeTrack1(time.Now(), "Main")
	s := "Go : Welcome to Go Programming"
	fmt.Println("worldCount : ", worldCount(s))
	fmt.Println("charCount : ", charCount(s))
}

func timeTrack1(start time.Time, name string) {
	end1 := time.Since(start)
	log.Printf("%s take %s", name, end1)
}

func worldCount(s string) map[string]int {
	ss := strings.Fields(s)
	ret := make(map[string]int)
	for i := 0; i < len(ss); i++ {
		ret[ss[i]]++
	}
	return ret
}

func charCount(s string) map[string]uint16 {
	//ss := strings.Fields(s)
	ret := make(map[string]uint16, len(s))
	for _, r := range s {
		ret[string(r)]++
	}
	return ret
}
