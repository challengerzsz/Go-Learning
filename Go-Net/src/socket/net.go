package main

import (
	"fmt"
	"net"
)

func GetConnection(protocol, addr string) (net.Conn, error) {

	conn, err := net.Dial(protocol, addr)
	conn.Close()
	return conn, err
}

func main() {
	conn, err := GetConnection("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
	}
	if conn != nil {
		conn.Close()
	}
}
