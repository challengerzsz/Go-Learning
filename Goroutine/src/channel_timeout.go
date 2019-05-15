package main

import (
	"fmt"
	"time"
)

func TimeoutResolver(ch chan int) {

	timeout := make(chan bool, 1)
	go func() {
		// 等待1s
		time.Sleep(1e9)
		timeout <- true
	}()

	select {
	case <-ch:
		// 处理ch中的数据
	case <-timeout:
		// 处理超时之后的timeout中的数据，此时ch没有收到任何数据，因为没有别的goroutine向这个ch中写入数据
		value, ok := <-timeout
		if ok {
			fmt.Println(value)
			return
		}
	}
}

func main() {

	ch := make(chan int, 1)
	TimeoutResolver(ch)

}
