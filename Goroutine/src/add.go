package main

import "fmt"

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	// 控制台不会打印任何值 因为go语言从初始package main
	// 并且执行main函数的时候不会等待别的goroutine的完成才完成 main函数执行完毕不会等待 会直接结束
	for i := 0; i < 10; i++ {
		go Add(i, i)
	}
}
