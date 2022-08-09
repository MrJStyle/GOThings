package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

func writeGOB(info map[string]string) {
	name := "demo.gob"
	File, _ := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0777)
	defer File.Close()
	encoder := gob.NewEncoder(File)
	if err := encoder.Encode(info); err != nil {
		fmt.Println(err)
	}
}

func readGOB(filePath string) {
	var M map[string]string
	File, _ := os.Open(filePath)
	decoder := gob.NewDecoder(File)
	decoder.Decode(&M)
	fmt.Println(M)

}

func main() {
	info := map[string]string{
		"name":    "golang",
		"website": "http://c.biancheng.net/golang/",
	}

	writeGOB(info)
	readGOB("demo.gob")
}
