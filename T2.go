package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	f, err := os.Create("./flag2.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	tpl := template.Must(template.ParseFiles("./T2.gohtml"))
	tpl.Execute(os.Stdout, slice)
	tpl.Execute(f, slice)
}
