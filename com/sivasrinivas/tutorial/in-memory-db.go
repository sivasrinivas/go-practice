package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {

	db := make(map[string]string, 1)

	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		log.Panic(err)
	}
	line = strings.TrimSpace(line)

	for strings.Compare(strings.ToUpper(line), "EXIT") != 0 {

		//KEYWORD KEY VALUE
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "SET":
			db[parts[1]] = parts[2]
		case "GET":
			log.Println(db[parts[1]])
		case "DEL":
			delete(db, parts[1])
		default:
			log.Println("Command not found")
		}

		line, err = r.ReadString('\n')
		if err != nil {
			log.Panic(err)
		}
		line = strings.TrimSpace(line)
	}

}
