package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle start!")
	var ctx context.Context
	ctx = r.Context()

	complete := make(chan struct{})

	go func() {
		// do something
		time.Sleep(5 * time.Second)
		complete <- struct{}{}
	}()

	select {
	case <-complete:
		fmt.Println("finish doing something")
	case <-ctx.Done():
		err := ctx.Err()
		if err != nil {
			fmt.Println("context deadline exceeded, cancel")
			fmt.Println(err.Error())
		}
	}
	fmt.Println("handler ends")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatalln(http.ListenAndServe(":8000", nil))
}
