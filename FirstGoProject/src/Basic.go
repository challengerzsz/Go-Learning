package main

import (
	"fmt"
	"math"
)


//            返回值列表
func GetName () (firstName, lastName, nickName string) {
	return "aaa", "bbb", "xxx"
}

//常量
const Pi float64 = 3.1415926
const (
	size int64 = 1024
	eof = -1
)
const ca, cb, cc = 3, 4, "foo"
//右值为编译期执行的位运算表达式
const mask = 1 << 3
//右值不能是运行期执行的表达式 下面语句会报错
//const home = os.Getenv("HOME")

const ( //iota遇见const会置0
	c0 = iota //0
	c1 = iota //1  每次再调用就会+1
	c2 = iota //2
)
const c3 = iota//0
const c4 = iota//0
const ( //iota遇见const会置0
	c5 = 1 << iota //1
	c6 //2  与上面初始化一样的话可以省略
	c7 //4
)
//go没有enum
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func testLeixing() {
	//bool
	var v1 bool
	v1 = true
	v2 := (false == v1)
	fmt.Println(v2)
	//bool不支持自动类型转换 也不支持强制类型转换
	//下面会编译错误
	//v1 = 1
	//v1 = bool(1)

	// int int32 是两种类型 不会自动类型转换
	var value32 int32
	value := 64 //自动推断为int
	//value = value32 //不会自动类型转换 编译错误
	value = int(value) //可以强制转换
	fmt.Println(value32, value)

	//不同类型的整形变量不能直接比较
	var i int32
	var j int64
	i, j = 1, 2

	//if i == j {   编译不通过 不同类型的整形变量不能直接比较
	//	fmt.Println(i, j)
	//}

	if i == 1 || j == 2{ //各种类型的整形变量可以直接与字面值比较
		fmt.Println(i, j)
	}

	//取反是^ c里面的取反是~
	fmt.Println(^i)

	//浮点数只有float32 float64
	var fnum1 float32
	fnum1 = 32
	fnum2 := 32.0//自动推断为float64

	//编译错误
	//fnum1 = fnum2
	fnum1 = float32(fnum2)//强制转换
	fmt.Println(fnum1)

	var vc1 complex64
	vc1 = 3.2 + 12i
	vc2 := 3.2 + 12i
	v3 := complex(3.2, 12)
	fmt.Println(vc1, vc2, v3)
	//分别输出实部虚部
	fmt.Println(real(v3), imag(v3))

}

//浮点数比较不精确
func IsEqual (f1, f2, p float64) bool {
	return math.Abs(f1 - f2) < p//自定义精确度
}


func testString() {
	//go中string也是基本类型
	var str string
	str = "hello world"
	ch := str[0]
	fmt.Println(len(str))
	fmt.Printf("%s, %c", str, ch)

	//编译错误 string不可变
	//str[0] = 'x'

	str = str + "jtt"
	ch = "hello"[0]

	//字符串遍历
	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i]  //用下标取出的类型是byte
		fmt.Println(i, ch) //输出为数字
	}

	for i, ch := range str {
		fmt.Println(i, ch)
	}


}

const N int = 200
func testArray() {
	array1 := [32]byte{}                                   //每个元素为一字节
	array2 := [2 * N] struct{ x, y int32 }{{1, 2}, {1, 2}} //复杂类型数组
	//array3 := [1000]*float64{}                             //指针数组
	//array4 := [2][2][2]float64{}                           //三维数组

	//遍历
	for i := 0; i < len(array1); i++ {
		fmt.Println(i, array1[i])
	}

	for i, v := range array2 {
		fmt.Println(i, v)
	}

	array := [5]int{1, 2, 3, 4, 5}
	modify(array)
	fmt.Println(array)
}

//参数是复制的 传值
func modify(array [5]int) {
	array[0] = 11
	fmt.Println(array)
}

func testSlice() {
	var array [10]int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//array[strat:end]切片
	var slice1 []int = array[:5]//前五个元素切片
	for _, v := range array {
		fmt.Println(v)
	}
	for _, v := range slice1 {
		fmt.Println(v)
	}

	//make()创建切片
	//slice2 := make([]int, 5) //初始元素个数5 初始值0
	//slice3 := make([]int, 5, 10)//预留10个元素空间
	slice4 := []int{1, 2, 3, 4, 5} //直接创建并初始化5个元素的切片
	for i, v := range slice4 {
		fmt.Println(i, v)
	}

	//切片会动态增减元素    分配的内存不够时会自动分配更大的一块内存 并把原来的元素复制到新内存
	//重新分配内存废时间 期初设置较大的空间可以空间换时间

	slice := make([]int, 5, 6)
	fmt.Println(len(slice))//长度
	fmt.Println(cap(slice))//分配的空间大小

	slice = append(slice, 1, 1, 1, 1)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))//两倍扩充

	slice = append(slice, slice4...) //省略号表示把slice打散添加 因为后面参数需要int
	newSlice := slice[:5]
	for i, v := range newSlice {
		fmt.Println(i, v)
	}

	slicea := []int{10, 11, 12, 13, 14}
	sliceb := []int{1, 2, 3, 4, 5, 6, 7}

	//copy(sliceb, slicea)//将a的所有元素复制到b
	//for _, v := range sliceb {
	//	fmt.Println(v)
	//}
	copy(slicea, sliceb)//将b前五个元素复制到a
	for _, v := range slicea {
		fmt.Println(v)
	}

}

type PersonInfo struct {
	ID int
	Name string
	Address string
}

func testMap() {
	//声明             键类型    值类型
	var personmap map[string] PersonInfo
	//make()创建
	personmap = make(map[string] PersonInfo)
	//添加数据
	personmap["aa"] = PersonInfo{1, "aa", "aaa"}
	personmap["bb"] = PersonInfo{2, "bb", "bbb"}
	personmap["cc"] = PersonInfo{3, "cc", "ccc"}

	person, ok := personmap["aa"]
	//fmt.Println(person.Address, ok)
	if ok {
		fmt.Println(person.Address)
	}

	//删除数据
	delete(personmap, "cc")

}

func testSentence() {
	x := testIfElse(5)
	fmt.Println(x)

}

//编译不报错了？
func testIfElse(x int) int {
	if x == 0 {
		return -1
	} else {
		return x
	}

}


func main() {
	var (
		v1 int
		v2 string
	)

	fmt.Println(v1, v2)

	var v3 int = 10
	var v4 = 10
	//自动推断类型 := 表示同时进行声明和初始化操作
	//使用:=的变量名不应该是重复的否则会出现错误
	//v4 := 9
	v5 := 10

	fmt.Println(v3, v4, v5)


	i := 1
	j := 2
	i, j = j, i
	fmt.Println("swap", i, j)

	//舍弃前两个返回值 不需要的返回值就不需要去定义了
	_, _, nickName := GetName()
	fmt.Println(nickName)

	//testString()
	//testArray()
	//testSlice()
	//testMap()
	testSentence()

}
