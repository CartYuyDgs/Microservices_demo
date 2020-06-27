package main

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
	"router",
}

type DirGenerator struct {
	dirList []string
}

func init() {
	//dir := &DirGenerator{
	//	dirList: ALLDirList,
	//}
	//Register("dirGenerator", dir)
}
