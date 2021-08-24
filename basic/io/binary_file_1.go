package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

func writeBinary(filePath string, info string) {
	file, _ := os.Create(filePath)

	defer file.Close()

	encoder := gob.NewEncoder(file)
	err := encoder.Encode(info)
	if err != nil {
		fmt.Println("fail to encode", err.Error())
		return
	} else {
		fmt.Println("encode successfully")
	}
}

func readBinary(filePath string) {
	file, _ := os.Open(filePath)

	defer file.Close()

	info := ""
	decoder := gob.NewDecoder(file)
	err := decoder.Decode(&info)
	if err != nil {
		fmt.Println("fail to decode", err.Error())
		return
	} else {
		fmt.Println("decode successfully")
		fmt.Println(info)
	}

}

func main() {
	filePath := "output.gob"
	info := "https://c.biancheng.net/golang/"

	writeBinary(filePath, info)
	readBinary(filePath)
}
