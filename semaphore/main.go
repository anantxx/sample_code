package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {
	defer timeTrack(time.Now(), "Main")
	var sem = semaphore.NewWeighted(int64(10))
	var ctx = context.Background()
	for i := 1; i < 100; i++ {
		fmt.Println("Semaphone Acquire : ", i)
		if err := sem.Acquire(ctx, 1); nil != err {
			fmt.Println("Semaphone Acquire Error", err.Error())
		}
		go func(j int) {
			remoteRPCCall(j)
			sem.Release(1)
		}(i)
	}
}

func remoteRPCCall(i int) {
	//fmt.Println("Start RemoteRPCCall : ", i)
	time.Sleep(time.Duration(int(i%1)) * time.Second)
	fmt.Println("End RemoteRPCCall : ", i)
}

func timeTrack(start time.Time, name string) {
	log.Printf("Start Time : %s", start)
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
