package main

import "fmt"

type Stringer interface {
	String() string
}

func Println(args ...interface{}) {
	for _, arg := range args {
		switch v := arg.(type) {
		case int:
			fmt.Println(v)
		case float64:
		default:
			if v, ok := arg.(Stringer); ok {
				val := v.String()
				fmt.Println(val)
				// ...
			} else {
				// ...
			}
		}
	}
}

func main() {
	var v1 interface{} = 1
	switch v := v1.(type) {
	case int:
		fmt.Println("int", v)
	case float32:
		fmt.Println("float32", v)
	}
}
