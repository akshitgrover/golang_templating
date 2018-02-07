package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	f, err := os.Create("flag1.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl := template.Must(template.ParseFiles("./T1.gohtml"))
	tpl.Execute(os.Stdout, "TEXT")
	tpl.Execute(f, "TEXT")
}
