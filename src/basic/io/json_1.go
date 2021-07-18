package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Website struct {
	Name   string `xml:"name,attr"`
	Url    string
	Course []string
}

func writeJSON(websites []Website) {
	// 创建文件
	filePtr, err := os.Create("info.json")
	if err != nil {
		fmt.Println("Fail to create info.json", err.Error())
		return
	}
	defer func(filePtr *os.File) {
		err := filePtr.Close()
		if err != nil {
			return
		}
	}(filePtr)

	// create json encoder
	encoder := json.NewEncoder(filePtr)

	err = encoder.Encode(websites)

	if err != nil {
		fmt.Println("error encoding", err.Error())
	} else {
		fmt.Println("encoding successfully")
	}
}

func readJSON(filePath string) {
	filePtr, err := os.Open(filePath)
	if err != nil {
		fmt.Println("fail to open file", filePath)
		return
	}

	defer func(filePtr *os.File) {
		err := filePtr.Close()
		if err != nil {
			return
		}
	}(filePtr)

	var websites []Website

	// craete json decoder
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&websites)
	if err != nil {
		fmt.Println("error decoding", err.Error())
		return
	} else {
		fmt.Println("decoding successfully")
		fmt.Println(websites)
	}

}

func main() {
	websites := []Website{
		{
			"Golang",
			"http://c.biancheng.net/golang/",
			[]string{"http://c.biancheng.net/cplus/", "http://c.biancheng.net/linux_tutorial/"},
		},
		{
			"Java",
			"http://c.biancheng.net/java/",
			[]string{"http://c.biancheng.net/socket/", "http://c.biancheng.net/python/"},
		},
	}

	writeJSON(websites)
	readJSON("info.json")
}
