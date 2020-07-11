package main

var main_template = `
package main

import (

	{{if not .Prefix}}
	"generate/{{.Package.Name}}"
	{{else}}
	{{.Package.Name}} "{{.Prefix}}/generate"
	{{end}}

	{{if not .Prefix}}
	"controller"
	{{else}}
	"{{.Prefix}}/router"
	{{end}}

	//{{if not .Prefix}}
	//"controller"
	//{{else}}
	//"{{.Prefix}}/controller"
	//{{end}}
	"google.golang.org/grpc"
	"log"
	"net"
)

var server = &router.RouterServer{}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("failed to listen: %v ", err)
	}
	s := grpc.NewServer()
	hello.Register{{.Service.Name}}Server(s, server)
	s.Serve(lis)
}
`
