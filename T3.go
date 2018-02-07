package main

import (
	"log"
	"os"
	"text/template"
)

type Person struct {
	Name   string
	Wealth int
}

func main() {
	f, err := os.Create("./flag3.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	struc := Person{
		Name:   "akshitgrover",
		Wealth: 1000000000,
	}
	tpl := template.Must(template.ParseFiles("./T3.gohtml"))
	tpl.Execute(f, struc)
	tpl.Execute(os.Stdout, struc)
}
