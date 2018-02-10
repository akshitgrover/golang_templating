package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func add(x []int) int {
	return x[0] + x[1]
}

func main() {
	fm := template.FuncMap{
		"add": add,
	}
	t := template.Must(((template.New("")).Funcs(fm)).ParseFiles("./T4.gohtml"))
	f, err := os.Create("flag4.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	slice := []int{1, 2}
	t.ExecuteTemplate(os.Stdout, "T4.gohtml", slice)
	t.ExecuteTemplate(f, "T4.gohtml", slice)
}
