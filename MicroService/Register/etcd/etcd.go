package etcd

import (
	"Microservices_demo/MicroService/Register"
	"github.com/etcd-io/etcd/clientv3"
)

type EtcdRegistry struct {
	options     *Register.Options
	client      *clientv3.Client
	serviceChan chan Register.Service
}
