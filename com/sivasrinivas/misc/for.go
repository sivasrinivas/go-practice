package main

import "fmt"

func main() {
	var i int = 1
	for i <= 3 { //for like while loop
		fmt.Println(i)
		i += 1
	}

	for j := 4; j <= 10; j++ { //Regular for loop
		fmt.Println(j)
	}

	for { //Infinite loop
		fmt.Println("infinite loop")
		break
	}
}
