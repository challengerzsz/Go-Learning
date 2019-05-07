package main

import "fmt"

func MyFunc(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "int")
		case string:
			fmt.Println(arg, "string")
		case int64:
			fmt.Println(arg, "int64")
		default:
			fmt.Println(arg, "unknown")
		}

	}
}

func main() {
	var (
		v1 int     = 1
		v2 int64   = 234
		v3 string  = "hello"
		v4 float32 = 1.234
	)

	MyFunc(v1, v2, v3, v4)

	// 匿名函数可直接赋值给一个变量或者直接执行
	// 函数体之后可以直接给参数进行执行
	f := func(a, b int, z float64) bool {
		v := int(z)
		fmt.Println("v : ", v)
		return a*b < int(z)
	}(1, 2, 2.2)

	fmt.Println("result ", f)

}
