package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	d := time.Now().Add(50 * time.Millisecond)
	//ctx, cancel := context.WithCancel(ctx)
	ctx, cancel := context.WithDeadline(ctx, d)
	//ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	go start(ctx, 5, "Hello World")
	//cancel()
	time.Sleep(10 * time.Second)
}

func start(ctx context.Context, timer int, message string) {
	select {
	case <-time.After(time.Duration(timer) * time.Second):
		fmt.Println(message)
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}
