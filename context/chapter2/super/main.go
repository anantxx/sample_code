package main

import (
	"fmt"
	"net/http"
	"time"

	"log"
)

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8081", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Printf("Super Handle Started")
	defer log.Printf("Super Handle Ended")

	select {
	case <-time.After(1 * time.Second):
		log.Print("SUCCESS")
		fmt.Fprintln(w, "Hello")
	case <-ctx.Done():
		err := ctx.Err()
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	time.Sleep(1 * time.Second)
}
