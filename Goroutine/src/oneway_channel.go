package main

import (
	"fmt"
)

// 声明这个函数的时候就希望这个channel是一个只能够读取int类型数据的channel
// 所以实现的时候除非进行了强制类型转换 否则将会只能进行读取 这样就实现了最小权限原则
func Parse(channel chan<- int) {
	for value := range channel {
		fmt.Println(value)
	}
}

func main() {

	// 单向channel
	// 只允许写入float64的channel
	var ch1 chan<- float64
	// wrong op 不能读取channel的数据
	//value := <-ch1

	// 只允许读取类型为int的channel
	var ch2 <-chan int
	// wrong op 不能向该channel进行写入数据
	//ch2 <- 2

	// ch3是一个可读可写的channel
	ch3 := make(chan int, 1)

	// ch4是一个单项读取的channel
	ch4 := <-chan int(ch3)

	// ch5是一个单项写入的channel
	ch5 := chan<- int(ch3)

	close(ch3)
	x, ok := <-ch3
	// ok这个返回值为false的时候表示这个channel被关闭了
	if ok == false {
		fmt.Println(" ch3 is closed !")
	}

	fmt.Println(ch1, ch2, ch4, ch5, x)
}
