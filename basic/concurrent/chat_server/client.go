package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)

	err = conn.Close()
	if err != nil {
		return
	}

	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	fmt.Println("copy")
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
