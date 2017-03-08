package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	tokens := strings.Split(input, ":")
	hour, _ := strconv.Atoi(tokens[0])
	min, _ := strconv.Atoi(tokens[1])
	sec, _ := strconv.Atoi(tokens[2][:len(tokens[2])-3])

	if hour == 12 && strings.Contains(tokens[2], "AM") {
		hour = 0
	} else if hour == 12 && strings.Contains(tokens[2], "PM") {
		//no op
	} else if strings.Contains(tokens[2], "PM") {
		hour += 12
	}
	fmt.Printf("%02d:%02d:%02d", hour, min, sec)
}
