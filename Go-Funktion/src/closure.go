package main

import "fmt"

// 闭包
func main() {

	var j int = 5
	a := func() func() {
		var i int = 10
		// 只有内部的匿名函数可以访问i
		return func() {
			fmt.Printf("i, j : %d %d \n", i, j)
		}
	}()

	a()

	j *= 2
	a()
}
