package main

import (
	"bytes"
	"fmt"
)

func main() {
	//scan := bufio.NewScanner(os.Stdin)
	//scan.Scan()
	//line := scan.Text()
	line := "ab"
	var buf bytes.Buffer
	i := 0
	for i < len(line) {
		if i+1 < len(line) {
			if line[i] == line[i+1] {
				i += 2
			} else {
				buf.WriteByte(line[i])
				i++
			}
		} else {
			buf.WriteByte(line[i])
			i++
		}
	}
	fmt.Println(buf.String())
}
