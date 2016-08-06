package main

import (
	"fmt"
	"time"
)

func f(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, i)
	}
}

func main() {
	go f("routine")
	time.Hour
	time.Sleep(time.Second*3)
	fmt.Println("Done")
}
