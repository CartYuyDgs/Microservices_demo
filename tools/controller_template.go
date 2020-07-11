package main

var controller_template = `
package controller

import (
	"context" 
	{{if not .Prefix}}
	"generate/{{.Package.Name}}"
	{{else}}
	{{.Package.Name}} "{{.Prefix}}/generate"
	{{end}}
) 

type {{.Rpc.Name}}Controller struct{}

func (s *{{.Rpc.Name}}Controller) CheckParams( ctx context.Context, r* {{$.Package.Name}}.{{.Rpc.RequestType}})(resp* {{$.Package.Name}}.{{.Rpc.ReturnsType}}, err error){ 
 return 
}

func (s *{{.Rpc.Name}}Controller) Run( ctx context.Context, r* {{$.Package.Name}}.{{.Rpc.RequestType}})(resp* {{$.Package.Name}}.{{.Rpc.ReturnsType}}, err error){ 
 return 
}
`
