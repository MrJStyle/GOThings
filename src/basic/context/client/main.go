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
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 6*time.Second)
	defer cancel()

	request, err := http.NewRequest(http.MethodGet, "http://localhost:8000", nil)
	if err != nil {
		log.Fatalln(err)
	}

	request = request.WithContext(ctx)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	fmt.Printf("%s", respBytes)
}
