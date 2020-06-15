package main

var main_template = `
package main

import (
	"Microservices_demo/tools/output/controller"
	hello "Microservices_demo/tools/output/generate"
	"google.golang.org/grpc"
	"log"
	"net"
)

var server = &controller.Server{}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("failed to listen: %v ", err)
	}
	s := grpc.NewServer()
	hello.RegisterHelloServiceServer(s, server)
	s.Serve(lis)
}
`
