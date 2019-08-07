package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d): // Continue execution
		fmt.Println(msg)
	case <-ctx.Done(): // This is for Cancel
		log.Println(ctx.Err())
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}()
	sleepAndTalk(ctx, 5*time.Second, "hello")
	defer cancel()
}
