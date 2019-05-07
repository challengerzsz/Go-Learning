package main

import (
	"errors"
	"fmt"
)

// 小写字母开头的函数只在本包内可见，大写字母开头的函数才可以被别的包所用

// 参数列表中的类型如果一致，可以省略前面的类型保留最后的类型
func Add(a, b int) (ret int, err error) {

	if a < 0 || b < 0 {
		err = errors.New("不支持负数")
		return
	}
	return a + b, nil
}

// 单返回值
func Decrease(a, b int) int {
	return 1
}

// 不定长参数必须处于函数的最后一个参数
// 实现原理上，可变长参数是一个数组切片，也就是[]type
func MultiArgs(name string, args ...int) {

	fmt.Println("name %s", name)
	for _, v := range args {
		fmt.Println("v: ", v)
	}

	// 可通过传递片段的形式传递给可变长参数，需要加上...
	//MultiArgs("zsz", args[:1]...)
}

func main() {

}
