package src

import "fmt"

func byteSlice(b []byte) []byte {
	return b
}

func runeSlice(r []rune) []rune {
	return r
}

func main() {

	b := []byte{0, 1}
	u8 := []uint8{2, 3}
	fmt.Printf("%T %T \n", b, u8)
	fmt.Println(byteSlice(b))
	fmt.Println(byteSlice(u8))

	r := []rune{4, 5}
	i32 := []int32{6, 7}
	fmt.Printf("%T %T \n", r, i32)
	fmt.Println(runeSlice(r))
	fmt.Println(runeSlice(i32))
}
