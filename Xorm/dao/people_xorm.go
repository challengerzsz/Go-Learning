package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type User struct {
	Id int64 `xorm:"'id'"`
	Username string `xorm:"'username'"`
	Age int `xorm:"'age'"`
	Sex int `xorm:"'sex'"`
	CreateTimeStamp string `xorm:"'create_ts'"`
}

const (
	DriverName = "mysql"
	MasterDataSourceName = "root:agytorudhcv11@tcp(localhost:3306)/xorm_db?charset=utf8"
)
func newEngine() *xorm.Engine {
	engine, err := xorm.NewEngine(DriverName, MasterDataSourceName)
	if err != nil {
		fmt.Print("new engine error!", err)
		return nil
	}
	engine.ShowSQL(true)
	err = engine.Sync2(new(User))
	if err != nil {
		fmt.Println("error")
	}
	return engine
}

func GetAllUser(user *User) bool {

	x, err := newEngine().Table("xorm_test_table").Where("id = ?", 10).Get(user)
	if err != nil {
		fmt.Println("select failed", err)
	}
	return x
}

func main() {
	var user User
	result := GetAllUser(&user)
	fmt.Println(result)
}
