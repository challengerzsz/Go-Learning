package src

import (
	"fmt"
	"io"
	"math"
	"os"
)

func GetName() (firstName, lastName, nickName string) {
	return "May", "zsz", "jtt"
}

func Equals(v1, v2, p float64) bool  {
	return math.Abs(v1 - v2) < p
}

func main() {
	var v1 int
	var v2 string
	// 数组
	var v3 [10]int
	// 数组切片
	var v4 []int
	var v5 struct {
		f int
	}
	var v6 *int
	// map key为string value为int
	var v7 map[string]int
	var v8 func(a int) int

	var (
		v9  int
		v10 string
	)

	// 变量初始化 后两种声明方式为编译器自动推导
	var v11 int = 10
	var v12 = 10
	// 同时进行变量声明和初始化
	v13 := 10

	// 交换 v11 和 v12
	v11, v12 = v12, v11

	// 忽略多重返回值其中的某几项
	_, _, nickName := GetName()

	// 常量定义
	const Pi float64 = 3.1415926
	// 无类型浮点常量
	const zero = 0.0
	const (
		size int64 = 1024
		eof        = -1
	)
	// 常量多重赋值
	const u, v float32 = 0, 3
	const a, b, c = 3, 4, "foo"
	// 可声明为一个编译器就确定下来的常量表达式
	const mask = 1 << 3
	// 常量声明为编译期确定不了的内容会出错
	//const Home = os.Getenv("HOME")

	// 预定义常量
	// iota 在每个const声明常量时被初始化为0
	const (
		c0 = iota
		c1 = iota
		c2 = iota
	)

	const (
		c3 = 1 << iota
		c4 = 1 << iota
		c5 = 1 << iota
	)

	const (
		u1         = iota * 42
		u2 float64 = iota * 42
		w          = iota * 42
	)

	const x = iota
	const y = iota

	// 如果const赋值语句一致 可省略
	const (
		c6 = iota // == 0
		c7        // == 1
		c8        // == 2
	)

	// 枚举
	const (
		Sunday = iota
		Monday
		// ...
	)

	// 布尔类型
	var v20 bool
	v20 = true
	v21 := (1 == 2)


	// int 和 int64属于两种不同类型 编译期会将互相赋值进行报错
	var int1 int64 = 100
	// int2会被自动确定为int
	int2 := 100
	int1 = int2
	// 使用强转可解决
	int1 = int64(int2)

	// 浮点型 float32 => float float63 => double
	var fvalue1 float32
	fvalue1 = 12
	// 自动推导时 如果不加小数点则会确定为类型int
	fvalue2 := 12.0


	// 复数类型
	var value4 complex64
	value4 = 3.2 + 12i
	value5 := 3.2 + 12i
	// complex 计算实部和虚部组成的复数
	value6 := complex(3.2, 12)

	// 使用函数求复数的实部和虚部
	result := real(value6)
	result1 := imag(value6)

	// 字符串
	var str string
	str = "hello world"
	ch := str[0]
	fmt.Printf("first letter is %c", ch)
	// 字符串内容不能再初始化之后被修改
	//str[0] = 'a'
	length := len(str)

	// 遍历数组
	for i := 0; i < length; i++ {
		fmt.Print(str[i])
	}

	// 以Unicode编码进行遍历
	for i, ch := range str {
		// ch类型为rune rune代表单个Unicode字符
		fmt.Println(i, ch)
	}

	// 数组
	const N = 1
	var array1 [32]byte
	// 复杂类型数组
	var array2 [2 * N] struct{x, y int32}
	// 指针数组
	var array3 [1000] *float64
	// 二维数组
	var array4 [3][5] int
	// 等同于[2] ([2] ([2] float64))
	var array5 [2][2][2] float64

}