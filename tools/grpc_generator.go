package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type GrpcGenerator struct {
}

func init() {
	grpc := &GrpcGenerator{}
	//log.Println("grpc init")
	Register("grpc generator", grpc)
}

func (g *GrpcGenerator) Run(opt *Option) (err error) {
	//D:\code\go\bin\proto  --go_out=plugins=grpc:. hello.proto
	//log.Println("grpc run")

	outPutParams := fmt.Sprintf("plugins=grpc:%s/generate/", opt.Output)

	cmd := exec.Command("D:\\code\\go\\bin\\protoc", "--go_out", outPutParams, opt.Proto3Filename)
	log.Println(cmd)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Println("cmd error :", err)
	}
	return
}
