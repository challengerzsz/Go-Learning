package main

import (
	"fmt"
	"myInterface"
	"os"
)

type Integer int

// 为指定类型添加方法 这个方法只有Integer类型的变量才可以进行调用
func (a Integer) Less(b Integer) bool {
	return a < b
}

// 当需要对一个引用变量进行修改的时候，需要通过传递指针这种形式进行修改
func (v *Integer) Add(b Integer) {
	*v += b
}

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

type Lesser interface {
	Less(b Integer) bool
}

type IStream interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
}

type Writer interface {
	Write(buf []byte) (n int, err error)
}

func main() {
	var m Integer = 1
	m.Less(2)
	// 只有Integer类型的变量能够使用Less函数 int类型无效
	var v int = 2
	fmt.Println(v)

	var a Integer = 5
	//a.Add(10)
	fmt.Println(a)

	// 将对象赋值给接口要求对象实现了接口的所有方法
	var c LessAdder = &a
	c.Add(10)
	fmt.Println(a)

	//var a1 Integer = 1
	//var b1 Lesser = a1

	var file1 IStream = new(os.File)
	var file2 myInterface.ReadWriter = file1
	var file3 IStream = file2
	fmt.Println(file3)

	// 因为Writer接口的方法集是IStream接口方法集的子集，所以可以赋值，相当于后者实现了前者
	var file4 Writer = file1
}
