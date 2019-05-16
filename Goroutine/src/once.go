package main

import (
	"fmt"
	"runtime"
	"sync"
)

var a string
// 借助Once实现全局唯一性操作
var Once sync.Once

func setup() {
	a = "hello world"
}

func doPrint() {
	// setup函数只会被执行一次
	// 会阻塞还未执行该函数的goroutine直到唯一的goroutine执行完该函数
	Once.Do(setup)
	fmt.Println(a)
}


func main() {
	go doPrint()
	runtime.Gosched()
	go doPrint()
	runtime.Gosched()
}
