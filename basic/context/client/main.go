package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	var ctx context.Context
	ctx = context.Background()
	var cancel func()
	ctx, cancel = context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var request *http.Request
	var err error
	request, err = http.NewRequest(http.MethodGet, "http://localhost:8000", nil)
	if err != nil {
		log.Fatalln(err)
	}

	request = request.WithContext(ctx)

	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var respBytes []byte
	respBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	fmt.Printf("%s", respBytes)
}
