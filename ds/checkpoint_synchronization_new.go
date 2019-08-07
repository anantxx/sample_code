package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func worker1(part string, c int, wg *sync.WaitGroup) {
	log.Println(part, "worker begins part ", c)
	time.Sleep(time.Duration(rand.Int63n(1e6)))
	log.Println(part, "worker completes part ", c)
	wg.Done()
}

var (
	partList1    = []string{"A", "B", "C", "D"}
	nAssemblies1 = 3
)

func main() {
	var wg, wg1 sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	wg1.Add(nAssemblies1)
	for c := 1; c <= nAssemblies1; c++ {
		go assembly(c, &wg, &wg1)
	}
	wg1.Wait()
	log.Println("assemble.  cycle")
}

func assembly(c int, wg, wg1 *sync.WaitGroup) {
	log.Println("begin assembly cycle", c)
	wg.Add(len(partList1))
	for _, part := range partList1 {
		go worker1(part, c, wg)
	}
	wg.Wait()
	wg1.Done()
}
