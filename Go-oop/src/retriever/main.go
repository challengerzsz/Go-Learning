package main

import (
	"fmt"
	"retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://www.baidu.com")
}

func main() {
	var r Retriever
	//r = mock.Retriever{"test.com"}
	r = real.Retriever{}
	fmt.Println(download(r))
}
