package main

import (
	"fmt"
	"html/template"
	"os"
)

type Persion struct {
	Name string
	Age  int
}

func main() {

	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Println("parse file err :", err)
		return
	}

	p := Persion{
		Name: "abc",
		Age:  19,
	}

	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("exec err :", err)
	}
}
