package main

import "github.com/emicklei/proto"

type Generator interface {
	Run(opt *Option, metaData *ServiceMateData) error
}

type ServiceMateData struct {
	service *proto.Service
	message []*proto.Message
	rpc     []*proto.RPC
}
