package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"log"
)

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Printf("Handle Started")
	defer log.Printf("Handle Ended")

	select {
	case <-time.After(1 * time.Second):
		strRes, ok := funHello(&ctx)
		if 1 == ok {
			fmt.Fprintln(w, strRes)
			log.Print("SUCCESS")
		}
	case <-ctx.Done():
		err := ctx.Err()
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//	time.Sleep(5 * time.Second)
}

func funHello(ctx *context.Context) (string, int) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8081", nil) // similar to res, err := http.Get("http://localhost:8080")

	if err != nil {
		log.Println("Error :", err)
		return "", 0
	}

	req = req.WithContext(*ctx)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error :", err)
		return "", 0
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)

	return bodyString, 1
	//	io.Copy(os.Stdout)
}
