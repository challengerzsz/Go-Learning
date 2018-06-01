package main

import (
	"net/http"
	"fmt"
	"log"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	//解析参数
	r.ParseForm()
	//输出到服务端的打印信息
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	//写入响应
	fmt.Fprint(w, "Hello test")

}


func main()  {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		 log.Fatal("ListenAndServe", err)
	}
}



