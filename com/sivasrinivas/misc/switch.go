package main

import "fmt"

func main() {
	var i int = 2
	fmt.Println("Variable i value is", i, ".")
	switch i {
	case 1:
		fmt.Print("One")
	case 2:
		fmt.Print("Two")
	case 3:
		fmt.Print("Three")
	default:
		fmt.Print("Default")
	}
}
