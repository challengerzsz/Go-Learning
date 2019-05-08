package main

import "fmt"

type Rect struct {
	x, y          int
	width, height int
}

func (rect *Rect) Add(k, v string) int {
	return rect.x + rect.y
}

// 构造函数 可以直接看出来new的过程 不像Java是不可见的new过程
func NewRect(x, y, width, height int) *Rect {
	return &Rect{x, y, width, height}
}

func main() {

	// 初始化一个对象
	rect1 := new(Rect)
	rect2 := &Rect{1, 2, 3, 4}
	rect3 := &Rect{}
	// 为进行初始化的变量都会被初始化为0值
	rect4 := &Rect{width: 10, x: 1}

	rect5 := NewRect(1, 2, 3, 4)
	fmt.Println(rect1, rect2, rect3, rect4, rect5)
}
