package main //packageWithValue returns a copy of parent in which the value associated with key is val. main

import (
	"context"
	"fmt"
)

func main() {
	type favContextKey1 string

	f := func(ctx context.Context, k favContextKey1) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey1("language")
	ctx := context.WithValue(context.Background(), k, "Go")
	f(ctx, k)

	ctx = context.WithValue(context.Background(), favContextKey1("color"), "Red")
	f(ctx, favContextKey1("color"))

	ctx = context.WithValue(context.Background(), favContextKey1("Animal"), "Dog")
	f(ctx, "Animal")

}
