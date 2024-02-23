package main

import (
	"fmt"
	"os"
)

func main() {
	str := "Hello, World!\nHello Golang!"
	wdata := []byte(str)

	f, err := os.Create("./test.txt")
	if err != nil {
		fmt.Println(err)
	}
	wcount, err := f.Write(wdata)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("write %d bytes:\n", wcount)

	rf, err := os.Open("./test.txt")
	defer rf.Close()
	defer func() {
		err := rf.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	rdata := make([]byte, 1024)
	rcount, err := rf.Read(rdata)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("read %d bytes:\n", rcount)
	fmt.Println(string(rdata[:rcount]))

}
