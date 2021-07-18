package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Website2 struct {
	Name   string `xml:"name,attr"`
	Url    string
	Course []string
}

func writeXML(websites Website2) {
	f, err := os.Create("info.xml")
	if err != nil {
		fmt.Println("fail to create info.xml")
		return
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			return
		}
	}(f)

	encoder := xml.NewEncoder(f)
	err = encoder.Encode(websites)
	if err != nil {
		fmt.Println("fail to encode")
	} else {
		fmt.Println("encode successfully")
	}
}

func readXML(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("fail to open", filePath)
		return
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			return
		}
	}(f)

	var websites Website2
	decoder := xml.NewDecoder(f)
	err = decoder.Decode(&websites)
	if err != nil {
		fmt.Println("fail to decode")
		return
	} else {
		fmt.Println("decode successfully")
		fmt.Println(websites)
	}

}

func main() {
	website := Website2{
		"Golang",
		"http://c.biancheng.net/golang/",
		[]string{"http://c.biancheng.net/cplus/", "http://c.biancheng.net/linux_tutorial/"},
	}

	writeXML(website)
	readXML("info.xml")
}
