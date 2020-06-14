package main

import (
	"fmt"
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

	fmt.Fprintf(file, "package main\n\n")

	fmt.Fprintf(file, "import (\n")

	fmt.Fprintf(file, "	\"Microservices_demo/tools/output/controller\" \n")
	fmt.Fprintf(file, `	hello "Microservices_demo/tools/output/generate"`)
	fmt.Fprintf(file, "\n")
	fmt.Fprintf(file, "	\"google.golang.org/grpc\"\n")
	fmt.Fprintf(file, "	\"log\" \n")
	fmt.Fprintf(file, "	\"net\" \n")

	fmt.Fprintf(file, "\n) \n\n")

	fmt.Fprintf(file, "var server = &controller.Server{}\n\n")

	fmt.Fprintf(file, "\n\n")

	fmt.Fprintf(file, `
func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("failed to listen: %v ", err)
	}
	s := grpc.NewServer()
	hello.RegisterHelloServiceServer(s, server)
	s.Serve(lis)
}
	`)
	return
}
