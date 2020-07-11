package main

import (
	"fmt"
	"github.com/emicklei/proto"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

var generatorMgr *GeneratorMgr = &GeneratorMgr{
	genMap:   make(map[string]Generator),
	metaData: &ServiceMateData{},
}

type GeneratorMgr struct {
	genMap   map[string]Generator
	metaData *ServiceMateData
}

func Register(name string, gen Generator) (err error) {
	_, ok := generatorMgr.genMap[name]
	if ok {
		err = fmt.Errorf("generator %s is exits", name)
		return
	}
	fmt.Println("name: ", name)
	generatorMgr.genMap[name] = gen
	return
}

func (g *GeneratorMgr) initOptout(opt *Option) (err error) {
	goPath := os.Getenv("GOPATH")
	if len(opt.Prefix) > 0 {
		if goPath[(len(goPath)-1)] != '/' {
			goPath = fmt.Sprintf("%s/src/", goPath)
		}
		opt.Output = path.Join(os.Getenv("GOPATH"), opt.Prefix)
		return
	}

	exeFilePath, err := filepath.Abs(os.Args[0])
	if err != nil {
		return
	}

	if runtime.GOOS == "windows" {
		exeFilePath = strings.Replace(exeFilePath, "\\", "/", -1)
		goPath = strings.Replace(goPath, "\\", "/", -1)
	}

	lastIdx := strings.LastIndex(exeFilePath, "/")
	if lastIdx < 0 {
		return
	}
	opt.Output = exeFilePath[0:lastIdx]

	opt.Output = path.Join(opt.Output, "/output/")
	//goPath = strings.ToLower(goPath)
	opt.Prefix = strings.Replace(opt.Output, goPath, "", -1)
	opt.Prefix = strings.Replace(opt.Prefix, "/src/", "", -1)
	log.Println(goPath)
	log.Println(opt.Output)
	log.Println(opt.Prefix)
	return
}

func (g *GeneratorMgr) Run(opt *Option) (err error) {
	err = g.initOptout(opt)
	if err != nil {
		return
	}

	err = g.parseService(opt)
	if err != nil {
		return
	}

	err = g.createAllDir(opt)
	if err != nil {
		return
	}

	g.metaData.Prefix = opt.Prefix
	for _, gen := range g.genMap {

		err := gen.Run(opt, g.metaData)
		if err != nil {
			return err
		}
	}
	return
}

func (g *GeneratorMgr) createAllDir(opt *Option) (err error) {
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

func (g *GeneratorMgr) parseService(opt *Option) (err error) {
	reader, err := os.Open(opt.Proto3Filename)
	if err != nil {
		log.Printf("openfile  %s failed, err %v", opt.Proto3Filename, err)
		return
	}

	defer reader.Close()

	parser := proto.NewParser(reader)
	definition, err := parser.Parse()
	if err != nil {
		log.Printf("parse file  %s failed, err %v", opt.Proto3Filename, err)
		return
	}

	proto.Walk(
		definition,
		proto.WithService(g.handleService),
		proto.WithMessage(g.handleMessage),
		proto.WithRPC(g.handleRpc),
		proto.WithPackage(g.handlePackage),
		proto.WithOption(g.handleOption),
	)

	//log.Println("parse proto success, rpc； ",c.rpc)
	//return c.generateRpc(opt)
	return
}

func (g *GeneratorMgr) handleService(s *proto.Service) {
	//fmt.Println(s.Name)
	g.metaData.Service = s
}

func (g *GeneratorMgr) handleMessage(m *proto.Message) {
	//fmt.Println(m.Name)
	g.metaData.Message = append(g.metaData.Message, m)
}

func (g *GeneratorMgr) handleRpc(r *proto.RPC) {
	g.metaData.Rpc = append(g.metaData.Rpc, r)
}

func (g *GeneratorMgr) handlePackage(r *proto.Package) {
	g.metaData.Package = r
	//g.metaData.Service.Name = g.metaData.Package.Name
}

func (g *GeneratorMgr) handleOption(r *proto.Option) {

}
