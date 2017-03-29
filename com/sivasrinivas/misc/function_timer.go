package main

import (
	"log"
	"time"
)

func StartTimer(name string) func() {
	start := time.Now()
	log.Println(name, "started")
	return func() {
		duration := time.Since(start)
		log.Println(name, "took", duration)
	}
}

func foo() {
	stop := StartTimer("foo")
	defer stop()
	time.Sleep(1 * time.Second)
}

func main() {
	foo()
}
