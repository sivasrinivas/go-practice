package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler1(w http.ResponseWriter, request *http.Request) {
	fmt.Fprint(w, "handler1")
}

func handler2(w http.ResponseWriter, request *http.Request) {
	fmt.Fprint(w, "handler2")
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/one/", handler1)
	router.HandleFunc("/two/", handler2)
	log.Print("Starting server...")
	http.ListenAndServe(":9999/", router)
}
