package main

import (
	"fmt"
	"github.com/emicklei/proto"
	"log"
	"os"
	"path"
	"text/template"
)

type ContralGenerator struct {
}

type RpcMeta struct {
	Rpc     *proto.RPC
	Package *proto.Package
	Prefix  string
}

func init() {
	c := &ContralGenerator{}
	Register("controller generator", c)
}

func (c *ContralGenerator) Run(opt *Option, mateData *ServiceMateData) (err error) {

	reader, err := os.Open(opt.Proto3Filename)
	if err != nil {
		log.Printf("openfile  %s failed, err %v", opt.Proto3Filename, err)
		return
	}

	defer reader.Close()
	return c.generateRpc(opt, mateData)
}

func (c *ContralGenerator) generateRpc(opt *Option, mateData *ServiceMateData) (err error) {

	for _, rpc := range mateData.Rpc {
		filename := path.Join("./", opt.Output, "controller", fmt.Sprintf("%s.go", rpc.Name))

		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			fmt.Printf("openfile %s failed, err: %v\n", filename, err)
			return err
		}

		defer file.Close()
		rpcMeta := &RpcMeta{}
		rpcMeta.Package = mateData.Package
		rpcMeta.Rpc = rpc
		rpcMeta.Prefix = mateData.Prefix
		err = c.render(file, controller_template, rpcMeta)
		if err != nil {
			fmt.Printf("render CONTROLLER failed, err:%v\n", err)
			return err
		}

	}

	return nil
}

func (c *ContralGenerator) render(file *os.File, data string, metaData *RpcMeta) (err error) {
	t := template.New("main")
	t, err = t.Parse(data)
	if err != nil {
		return
	}

	err = t.Execute(file, metaData)
	return
}
