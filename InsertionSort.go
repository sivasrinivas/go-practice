package main

import (
 "fmt"
 "bytes"
)

func main(){
	var buffer bytes.Buffer
	for i := 0; i<100; i++ {
		buffer.WriteString("a")
		buffer.WriteInt(i)
	}
	fmt.Println(buffer.String())
	fmt.Print("Min of 2,3: ")
}
