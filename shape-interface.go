package main

import (
	"fmt"
)

type Rectangle struct {
	length, width int
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

type Square struct {
	side int
}

func (s Square) Area() int {
	return s.side * s.side
}

type Shaper interface {
	Area() int
}

func main() {
	
	r := Rectangle{length:5, width:3}
	s := Square{5}
	var sh Shaper

	sh = r
	fmt.Println("Rectangle details: ", r)
	fmt.Println("Rectangle area: ", sh.Area())

	sh = s
	fmt.Println("Square details: ", s)
	fmt.Println("Square area: ", sh.Area())
}
