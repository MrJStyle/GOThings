package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func writeTxt(filePath string) {
	file, _ := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
	defer file.Close()

	//写入内容
	str := "http://c.biancheng.net/golang/\n" // \n\r表示换行  txt文件要看到换行效果要用 \r\n

	// 写入时，使用带缓存的*writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 3; i++ {
		_, err := writer.WriteString(str)
		if err != nil {
			return
		}
	}
	writer.Flush()
}

func readTxt(filePath string) {
	file, _ := os.Open(filePath)

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}
	fmt.Println("读取文件结束")
}

func main() {
	filePath := "out.txt"

	writeTxt(filePath)
	readTxt(filePath)
}
