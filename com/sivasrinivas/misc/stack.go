package main

import (
	"bufio"
	"fmt"
	"os"
)

func push(s []int, x int) []int {
	return append(s, x)
}

func pop(s []int) ([]int, int) {
	x := s[len(s)]
	return s[:len(s)-1], x
}

func main() {
	var stack []int

	for i := 0; i < 10; i++ {
		push(stack, i)
	}

	for i := 0; i < 10; i++ {
		_, x := pop(stack)
		fmt.Println(x)
	}
	reader := bufio.NewReader(os.Stdin)
	s, e := reader.Peek(2)
	e.Error()
	fmt.Println(s)
}
