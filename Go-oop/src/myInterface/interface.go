package main

import "fmt"

type Person struct {
	Name string
	age  int
}

type Float float64

func (f *Float) Add() {
	*f++
}

func (f *Float) Decrease() {
	fmt.Println(f)
	*f--
}

func (f Float) Decrease1() {
	fmt.Println(&f)
	(&f).Decrease()
}

func (person Person) getName(p Person) string {
	return p.Name
}

func (person *Person) addAge(p Person) bool {
	p.age++
	return true
}

type IPersonInterface interface {
	getName(p Person) string
	addAge(p Person) bool
}

type IFloatInterface interface {
	Add()
	Decrease()
}

type ReadWriter interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
}

type Reader struct {
}

func (r *Reader) Write(buf []byte) (n int, err error) {
	return 0, nil
}

func (r *Reader) Read(buf []byte) (n int, err error) {
	return 0, nil
}

func main() {

	// 接口赋值的时候需要明确到底是原始类型实现了接口还是指针类型实现了接口 这一点确定了赋值的时候是传递地址还是原始变量

	//p := &Person{"zsz", 21}
	//fmt.Println(p)
	// p本来就是一个指针
	//var personInterface IPersonInterface = p
	//fmt.Println(personInterface.getName(*p))
	//var v Float = 2.0
	//var floatInterface IFloatInterface = &v
	//floatInterface.Add()
	//fmt.Println(v)
	//floatInterface.Decrease()
	//fmt.Println(v)
	//fmt.Println(&v)
	//v.Decrease1()
	//fmt.Println(v)
	//
	//p1 := Person{"zsz", 21}
	//fmt.Println(p1)
	//
	reader := &Reader{}
	var readerWriter ReadWriter = reader

	// 接口查询需要等到运行阶段才会完成 接口赋值在编译阶段就能够确定
	// 查询ReadWriter接口指向的对象实例是否实现了某一接口
	// r和ok初始化 最后判断ok
	if r, ok := readerWriter.(ReadWriter); ok {
		fmt.Println(r, "ok!")
	}
}
