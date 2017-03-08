package main

import (
	"encoding/json"
	"fmt"
)

type ColorGroup struct {
	Id     int
	Name   string
	Colors []string
}

type ColorGroupJson struct {
	Id     int      `json:"id"`
	Name   string   `json:"name"`
	Colors []string `json:"colors"`
}

func main() {
	cGroup := ColorGroup{1, "siva", []string{"red", "blue", "green"}}
	jStr1, err := json.MarshalIndent(cGroup, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jStr1))

	cGroupJson := ColorGroupJson{1, "siva", []string{"red", "blue", "green"}}
	jStr2, err := json.MarshalIndent(cGroupJson, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jStr2))
}
