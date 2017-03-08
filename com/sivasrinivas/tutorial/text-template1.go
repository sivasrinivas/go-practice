package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

func main() {
	//Parse one file
	tpl, err := template.ParseFiles("src/go-practice/com/sivasrinivas/tutorial/file-one.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	//Add another file to existing template pointer
	tpl2, err := tpl.ParseFiles("src/go-practice/com/sivasrinivas/tutorial/file-two.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(os.Stdout, nil)
	tpl2.ExecuteTemplate(os.Stdout, "file-two.gohtml", nil)

	//Parse folder
	ftpl, err := template.ParseGlob("src/go-practice/com/sivasrinivas/tutorial/*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("\nGlob")
	ftpl.ExecuteTemplate(os.Stdout, "file-one.gohtml", nil)
	ftpl.ExecuteTemplate(os.Stdout, "file-two.gohtml", nil)
}
