package src

import (
	"fmt"
	"math"
)

func GetName() (firstName, lastName, nickName string) {
	return "May", "zsz", "jtt"
}

func Equals(v1, v2, p float64) bool {
	return math.Abs(v1-v2) < p
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
	var array2 [2 * N]struct{ x, y int32 }
	// 指针数组
	var array3 [1000]*float64
	// 二维数组
	var array4 [3][5]int
	// 等同于[2] ([2] ([2] float64))
	var array5 [2][2][2]float64

	// 定义数组切片
	// 基于数组
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 基于数组前5个元素创建的切片
	var mySlice []int = myArray[:5]
	fmt.Println("Elements of myArray")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}

	fmt.Println("Elements of mySlice")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}

	// 直接创建切片
	// 创建元素个数为5个的切片，元素初始值为0
	mySlice1 := make([]int, 5)
	// 创建初始元素为5个的切片，元素初始值为0，并预留10个元素的存储空间
	mySlice2 := make([]int, 5, 10)
	// 直接创建切片，会创建一个匿名数组
	mySlice3 := []int{1, 2, 3, 4, 5}

	// 10
	sliceAllocateLen := cap(mySlice2)
	// 5
	sliceLen := len(mySlice2)

	// append第二个参数是可变长参数
	mySlice := append(mySlice, 1, 2, 3, 4, 5)

	// 第二个参数是可变长参数所以加上...
	mySlice := append(mySlice, mySlice2...)

	// 基于数组切片创造数组切片
	oldSlice := []int{1, 2, 3, 4, 5}
	newSlice := oldSlice[:3]
	// 同样可以用以下写法，只要不超过cap()函数返回值，都可以创建，使用0补全
	newSlice := oldSlice[:10]

	// 切片内容复制
	slice4 := []int{1, 2, 3, 4, 5}
	slice5 := []int{5, 6, 7}
	// 将slice5的3个元素复制到slice4的前三个元素位置
	copy(slice4, slice5)
	// 将slice4的前3个元素复制给slice5
	copy(slice5, slice4)

}
