package main

import (
	"log"
	"os"
	"path"
)

var ALLDirList []string = []string{
	"controller",
	"idl",
	"main",
	"scripts",
	"conf",
	"app/router",
	"app/config",
	"model",
	"generate",
}

type DirGenerator struct {
	dirList []string
}

func (d *DirGenerator) Run(opt *Option, mateData *ServiceMateData) (err error) {
	for _, dir := range ALLDirList {
		fullDir := path.Join(opt.Output, dir)
		err := os.MkdirAll(fullDir, 0755)
		if err != nil {
			log.Println("mkdir dir err:", err)
			return err
		}
	}
	return
}

func init() {
	//dir := &DirGenerator{
	//	dirList: ALLDirList,
	//}
	//Register("dirGenerator", dir)
}
