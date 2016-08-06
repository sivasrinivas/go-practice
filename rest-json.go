package main

import (
  "fmt"
  "net/http"
  "encoding/json"
)

type Payload struct {
  Fruits Fruits
}

type Fruits map[string]int

func main() {
  fmt.Println("listening on port 3000...");
  http.HandleFunc("/", servRest)
  http.ListenAndServe("localhost:3000", nil)
}

func servRest(w http.ResponseWriter, r *http.Request) {
  response, err := createJson()
  if err!= nil {
    panic(err)
  }
  fmt.Fprintf(w, string(response))
}

func createJson() ([]byte, error) {
  fruits := make(map[string]int)
  fruits["apples"] = 25
  fruits["oranges"] = 10

  p := Payload{fruits}

  return json.Marshal(p)
}
