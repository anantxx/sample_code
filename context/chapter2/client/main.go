package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()
	//ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}()

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil) // similar to res, err := http.Get("http://localhost:8080")

	if err != nil {
		log.Fatal("Error :", err)
	}

	req = req.WithContext(ctx)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error :", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}

	io.Copy(os.Stdout, res.Body)
}
