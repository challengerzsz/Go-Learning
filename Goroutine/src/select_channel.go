package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	for {
		// select开始一个新的选择块 类似switch
		select {
		case ch <- 0:
		case ch <- 1:
		}
		i := <-ch
		fmt.Println("Value received : ", i)
	}

	// ch1被声明为了1024大小的channel这个时候就算没有别的goroutine进行读取这个channel也不会阻塞接下来的写入操作
	// 也叫这种channel为带缓冲的channel
	ch1 := make(chan string, 1024)
	fmt.Println(ch1)
}
