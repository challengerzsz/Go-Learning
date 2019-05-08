package main

import "fmt"

type Base struct {
	Name string
}

func (base *Base) Foo() {

}

func (base *Base) Bar() {
	fmt.Println("Base@Bar()")
}

type Foo struct {
	// 匿名组合
	Base
	// ...
}

// 通过Foo组合Base的方式 重写Base@Bar方法
func (foo *Foo) Bar() {
	fmt.Println("Foo@Bar()")
}

func main() {
	base := &Base{}
	foo := &Foo{}
	base.Bar()
	foo.Bar()
	foo.Base.Bar()
}
