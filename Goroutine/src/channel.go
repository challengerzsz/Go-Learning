package main

import "fmt"

func Count1(ch chan int) {
	fmt.Println("Counting")
	// 在这个channel被其他goroutine读取之前这个写入操作是阻塞的
	ch <- 1
}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count1(chs[i])
	}

	for _, ch := range chs {
		// 在这个channel被写入数据之前读取操作也是阻塞的
		<-ch
	}
}
