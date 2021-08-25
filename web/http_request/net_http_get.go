package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

const (
	binGetUrl = "http://httpbin.org/get"
)

func httpGet() {
	resp, err := http.Get(binGetUrl)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		fmt.Println("OK")
	}
}

// httpGetWithUrlParam request with URL parameters
func httpGetWithUrlParam() {
	params := url.Values{}
	Url, err := url.Parse(binGetUrl)
	if err != nil {
		return
	}
	params.Set("name", "zhaofan")
	params.Set("age", "23")
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println(urlPath)
	resp, err := http.Get(urlPath)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

type result struct {
	Args    string            `json:"args"`
	Headers map[string]string `json:"headers"`
	Origin  string            `json:"origin"`
	Url     string            `json:"url"`
}

// httpGetWithJsonResponse
// bind response body to a json
func httpGetWithJsonResponse() {
	resp, err := http.Get(binGetUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	var res result
	_ = json.Unmarshal(body, &res)
	fmt.Printf("%#v", res)
}

//
func httpGetRequestWithHeader() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", binGetUrl, nil)
	req.Header.Add("name", "Peter Park")
	req.Header.Add("age", "3")
	resp, _ := client.Do(req)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main() {
	//httpGet()
	//httpGetWithUrlParam()
	//httpGetWithJsonResponse()
	httpGetRequestWithHeader()
}
