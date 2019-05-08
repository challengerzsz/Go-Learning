package main

import "log"

type Job struct {
	// 匿名组合类型相当于是去掉包名的类型名称 在同一个结构体内重复会编译错误 也有可能不会编译错误 直到使用这两个同名的匿名组合类型时
	//*Logger
	Command string
	// 匿名组合
	*log.Logger
}

type Logger struct {
	Name string
}

func (job *Job) Start() {
	job.Println("starting now")
	// 做一些事情
	job.Println("started")
}

func main() {

}
