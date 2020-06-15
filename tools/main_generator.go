package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path"
)

type MainGenerator struct {
}

func init() {
	main := &MainGenerator{}
	Register("main generator", main)
}

func (g *MainGenerator) Run(opt *Option) (err error) {
	filename := path.Join(opt.Output, "main", fmt.Sprintf("main.go"))
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Println("open file %s failed, err %v", filename, err)
		return
	}

	defer file.Close()
	err = g.render(file, main_template)
	if err != nil {
		fmt.Println("render failed , err %v\n", err)
		return
	}
	return
}

func (g *MainGenerator) render(file *os.File, data string) (err error) {

	temp := template.New("main")
	t, err := temp.Parse(data)
	if err != nil {
		return
	}
	t.Execute(file, nil)
	return
}
