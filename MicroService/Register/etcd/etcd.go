package etcd

import (
	"Microservices_demo/MicroService/Register"
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"path"
)

type EtcdRegistry struct {
	options     *Register.Options
	client      *clientv3.Client
	serviceChan chan *Register.Service
}

var (
	etcdRegistry *EtcdRegistry = &EtcdRegistry{
		serviceChan: make(chan *Register.Service, 8),
	}
)

func init() {
	Register.RegisterPlugin(etcdRegistry)
	go etcdRegistry.run()
}

func (e *EtcdRegistry) Init(ctx context.Context, opts ...Register.Option) (err error) {
	e.options = &Register.Options{}
	for _, opt := range opts {
		opt(e.options)
	}

	e.client, err = clientv3.New(clientv3.Config{
		Endpoints:   e.options.Addrs,
		DialTimeout: e.options.TimeOut,
	})

	if err != nil {
		return err
	}
	return err
}

func (e *EtcdRegistry) Name() string {
	return "etcd"
}

func (e *EtcdRegistry) Register(ctx context.Context, service *Register.Service) (err error) {
	select {
	case e.serviceChan <- service:
	default:
		err = fmt.Errorf("register chan is full")
		return
	}
	return
}

func (e *EtcdRegistry) Unregister(ctx context.Context, service *Register.Service) (err error) {
	return
}

func (e *EtcdRegistry) run() {

}

func (e *EtcdRegistry) servicePath(service *Register.Service) string {
	return path.Join(e.options.RegistryPath, service.Name)
}
