package src

import (
	"fmt"
	"os"
	"syscall"
)

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

func Stat(name string) (fi os.FileInfo, err error) {
	var stat syscall.Stat_t

	err = syscall.Stat(name, &stat)

	if err != nil {
		return nil, &PathError{"stat", name, err}
	}

	return fileInfoFromStat(&stat, name), nil
}

func fileInfoFromStat(t *syscall.Stat_t, s string) os.FileInfo {
	// 省略...
	return nil
}

func main() {
	fi, err := Stat("a.txt")
	if err != nil {
		if e, ok := err.(*os.PathError); ok && e.Err != nil {
			// 获取PathError类型变量e中的其他信息并处理
			fmt.Println(fi.Name())
		}
	}
}
