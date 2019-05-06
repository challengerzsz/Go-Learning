package src

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)
	// 初始化map大小
	//personDB = make(map[string] PersonInfo, 100)
	personDB["1"] = PersonInfo{"1", "zsz", "xiyou"}
	personDB["2"] = PersonInfo{"2", "jtt", "xiyou"}

	person, ok := personDB["1"]
	if ok {
		fmt.Println("has found", person.Name)
	} else {
		fmt.Println("not found")
	}
	delete(personDB, "1")
}
