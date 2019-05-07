package src

import (
	"io"
	"log"
	"os"
)

// defer是栈式调用 最后一个defer先被执行
func CopyFile(dst, src string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}

	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return
	}

	defer dstFile.Close()

	// defer声明一个func 可以写很多的出现异常时关闭资源等等的动作 无论是否发生异常这个匿名函数都会被执行
	defer func() {

		//srcFile.Close()
		//dstFile.Close()

		if r := recover(); r != nil {
			log.Printf("Runtime error caught: %v", r)
		}
	}()

	return io.Copy(dstFile, srcFile)
}
