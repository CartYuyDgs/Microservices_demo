package main

import (
	"fmt"
	"github.com/emicklei/proto"
	"log"
	"os"
	"path"
)

type ContralGenerator struct {
	Service *proto.Service
	Mesage  []*proto.Message
	Rpc     []*proto.RPC
}

func init() {
	c := &ContralGenerator{}
	Register("controller generator", c)
}

func (c *ContralGenerator) Run(opt *Option, mateData *ServiceMateData) (err error) {

	//reader, err := os.Open(opt.Proto3Filename)
	//if err != nil {
	//	log.Printf("openfile  %s failed, err %v", opt.Proto3Filename, err)
	//	return
	//}
	//
	//defer reader.Close()
	//
	//parser := proto.NewParser(reader)
	//definition, err := parser.Parse()
	//if err != nil {
	//	log.Printf("parse file  %s failed, err %v", opt.Proto3Filename, err)
	//	return
	//}
	//
	//proto.Walk(
	//	definition,
	//	proto.WithService(c.handleService),
	//	proto.WithMessage(c.handleMessage),
	//	proto.WithRPC(c.handleRpc))
	//
	////log.Println("parse proto success, rpc； ",c.rpc)
	//return c.generateRpc(opt)
	return
}

func (c *ContralGenerator) generateRpc(opt *Option) (err error) {

	filename := path.Join(opt.Output, "controller", fmt.Sprintf("%s.go", c.Service.Name))
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Println("open file %s failed, err %v", filename, err)
		return
	}

	defer file.Close()
	fmt.Fprintf(file, "package controller\n\n")

	fmt.Fprintf(file, "import (\n")

	fmt.Fprintf(file, `	hello "Microservices_demo/tools/output/generate"`)
	fmt.Fprintf(file, "\n")
	fmt.Fprintf(file, "	\"context\" \n")
	fmt.Fprintf(file, "\n) \n\n")

	fmt.Fprintf(file, "type Server struct{}\n\n")
	for _, method := range c.Rpc {
		fmt.Fprintf(file, "func (s *Server) %s("+
			" ctx context.Context, r *hello.%s)(resp *hello.%s, err error){ \n return \n}\n\n", method.Name,
			method.RequestType, method.ReturnsType)
	}
	return
}

//func (c *ContralGenerator) handleService(s *proto.Service) {
//	//fmt.Println(s.Name)
//	c.service = s
//}
//
//func (c *ContralGenerator) handleMessage(m *proto.Message) {
//	//fmt.Println(m.Name)
//	c.mesage = append(c.mesage, m)
//}
//
//func (c *ContralGenerator) handleRpc(r *proto.RPC) {
//	//fmt.Println(r.Name)
//	//fmt.Println(r.RequestType)
//	//fmt.Println(r.ReturnsType)
//	//fmt.Println(r.Comment)
//
//	c.rpc = append(c.rpc, r)
//}
