package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func worker(part string, wg *sync.WaitGroup) {
	log.Println(part, "worker begins part")
	time.Sleep(time.Duration(rand.Int63n(1e6)))
	log.Println(part, "worker completes part")
	wg.Done()
}

var (
	partList    = []string{"A", "B", "C", "D"}
	nAssemblies = 3
)

func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	for c := 1; c <= nAssemblies; c++ {
		log.Println("begin assembly cycle", c)
		wg.Add(len(partList))
		for _, part := range partList {
			go worker(part, &wg)
		}
		wg.Wait()
		log.Println("assemble.  cycle", c, "complete")
	}
}
