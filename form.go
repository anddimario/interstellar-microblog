package main

import (
	"html/template"
	"os"
)

type Person struct {
    UserName string
}

func main() {
	t, _ := template.ParseFiles("tmpl/form.html")
	t.Execute(os.Stdout, nil)
}
