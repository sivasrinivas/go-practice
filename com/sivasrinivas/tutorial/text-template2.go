package main

import (
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) > 3 {
		s = s[0:3]
	}
	return s
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("src/go-practice/com/sivasrinivas/tutorial/file-three.gohtml"))
}

func main() {
	sagesMap := map[string]string{
		"Gandhi":  "India",
		"Buddha":  "India",
		"Lincoln": "America",
		"Mandela": "Africa",
	}

	sagesSlice := []string{"Gandhi", "Buddha", "Lincoln", "Mandela"}

	data := struct {
		SagesMap   map[string]string
		SagesSlice []string
	}{
		sagesMap,
		sagesSlice,
	}

	tpl.ExecuteTemplate(os.Stdout, "file-three.gohtml", data)
}
