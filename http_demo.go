package main

import (
	"fmt"
	"net/http"
	"time"
)

func client()  {
	resp, _ := http.Get("http://example.com/")

	fmt.Println(resp.Status)

	// 这些地儿为啥必须取地址？
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	fmt.Printf("%T \n", tr)
	client := http.Client{Transport: tr}
	resp, _ = client.Get("https://example.com")

	fmt.Print(resp.Status)
}




func main() {

	client()
}
