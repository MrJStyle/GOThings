package main

import "io"

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type device struct{}

func (d *device) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (d *device) Close() error {
	return nil
}

func main() {
	var wc WriterCloser = new(device)

	// write data
	wc.Write(nil)

	wc.Close()

	// 声明写入器, 并赋予device的新实例
	var writeOnly io.Writer = new(device)
	// 写入数据
	writeOnly.Write(nil)
}
