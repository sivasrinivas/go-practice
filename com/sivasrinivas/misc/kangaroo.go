package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	flag := false
	var location1, location2 int
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	tokens := strings.Split(input, " ")
	x1, _ := strconv.Atoi(tokens[0])
	v1, _ := strconv.Atoi(tokens[1])
	x2, _ := strconv.Atoi(tokens[2])
	v2, _ := strconv.Atoi(tokens[3])

	if x1 == x2 {
		flag = true
	} else {
		location1 = x1 + v1
		location2 = x2 + v2
		if location1 == location2 {
			flag = true
		} else if (v1 < v2 && location1 > location2) || (v1 > v2 && location1 < location2) {
			for i := 2; i <= 10000; i++ {
				location1 = location1 + v1
				location2 = location2 + v2
				if location1 == location2 {
					flag = true
					break
				}
			}
		}
	}

	if flag {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
