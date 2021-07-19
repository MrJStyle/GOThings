package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

type Website3 struct {
	Url int32
}

// write custom binary file

func writeCustomBinary(filePath string) {
	file, _ := os.Create(filePath)

	defer file.Close()

	for i := 0; i < 10; i++ {
		info := Website3{Url: int32(i)}

		var bin_buf bytes.Buffer

		binary.Write(&bin_buf, binary.LittleEndian, info)
		b := bin_buf.Bytes()
		_, err := file.Write(b)
		if err != nil {
			fmt.Println("fail to encode", err.Error())
			return
		}
	}
	fmt.Println("encode successfully")
}

func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		fmt.Println("fail to decode", err.Error())
	}
	return bytes
}

func readCustomBinary(filePath string) {
	file, _ := os.Open(filePath)

	defer file.Close()

	var websites []Website3
	var website = Website3{}

	for i := 1; i <= 10; i++ {
		data := readNextBytes(file, 4)
		buffer := bytes.NewBuffer(data)
		err := binary.Read(buffer, binary.LittleEndian, &website)
		if err != nil {
			fmt.Println("二进制文件读取失败", err)
			return
		}
		fmt.Println("第", i, "个值为：", website)
		websites = append(websites, website)
	}
	fmt.Println(websites)
}

func main() {
	filePath := "output.bin"

	writeCustomBinary(filePath)
	readCustomBinary(filePath)
}
