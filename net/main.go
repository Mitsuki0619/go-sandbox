package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}

	conn, err := ln.Accept()
	if err != nil {
		fmt.Println(err)
	}

	str := "Hello from Server"
	data := []byte(str)


	_, err = conn.Write(data)
	if err != nil {
		fmt.Println(err)
	}

	conn2, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
	}

	data2 := make([]byte, 1024)
	count, _ := conn2.Read(data)
	fmt.Println(string(data2[:count]))
	
}
