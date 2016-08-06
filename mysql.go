package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row := db.QueryRow("Select * from book where id=1")
	var id int
	var title string
	var col1 string
	row.Scan(&id, &title, &col1)

	fmt.Fprintf(w, "Book detials are: %d - %s - %s", id, title, col1)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
