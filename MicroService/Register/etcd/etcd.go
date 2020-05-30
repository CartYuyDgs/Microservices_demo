package etcd

import (
	"Microservices_demo/MicroService/Register"
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"log"
	"path"
	"time"
)

const MaxServiceNum = 8

type EtcdRegistry struct {
	options     *Register.Options
	client      *clientv3.Client
	serviceChan chan *Register.Service

	registryServiceMap map[string]*RegisterService
}

type RegisterService struct {
	id          clientv3.LeaseID
	service     *Register.Service
	registered  bool
	keepAliveCh <-chan *clientv3.LeaseKeepAliveResponse
}

var (
	etcdRegistry *EtcdRegistry = &EtcdRegistry{
		serviceChan:        make(chan *Register.Service, MaxServiceNum),
		registryServiceMap: make(map[string]*RegisterService, MaxServiceNum),
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
	for {
		select {
		case service := <-e.serviceChan:
			_, ok := e.registryServiceMap[service.Name]
			if ok {
				break
			}
			registryService := &RegisterService{
				service: service,
			}
			e.registryServiceMap[service.Name] = registryService
		default:
			e.registryOrKeepAlive()
			time.Sleep(time.Millisecond * 500)

		}
	}

}

func (e *EtcdRegistry) registryOrKeepAlive() {
	for _, registryService := range e.registryServiceMap {
		if registryService.registered {
			e.keepAlive(registryService)
			continue
		}

		e.registerService(registryService)
	}
}

func (e *EtcdRegistry) keepAlive(registryService *RegisterService) {
	select {
	case resp := <-registryService.keepAliveCh:
		if resp == nil {
			registryService.registered = false
			return
		}
		log.Printf("service:%s node:%s port:%v", registryService.service.Name,
			registryService.service.Nodes[0].Ip, registryService.service.Nodes[0].Port)
	}
	return
}

func (e *EtcdRegistry) registerService(registryService *RegisterService) {
	resp, err := e.client.Grant(context.TODO(), e.options.HeartBeat)
	if err != nil {
		log.Fatal(err)
	}
	registryService.id = resp.ID
	for _, node := range registryService.service.Nodes {
		tmp := &Register.Service{
			Name: registryService.service.Name,
			Nodes: []*Register.Node{
				node,
			},
		}

		data, err := json.Marshal(tmp)
		if err != nil {
			continue
		}
		key := e.serviceNodePath(tmp)
		log.Printf("register Key:%s\n", key)
		_, err = e.client.Put(context.TODO(), key, string(data), clientv3.WithLease(resp.ID))
		if err != nil {
			continue
		}

		ch, err := e.client.KeepAlive(context.TODO(), resp.ID)
		if err != nil {
			continue
		}

		registryService.keepAliveCh = ch

		registryService.registered = true
	}

}

func (e *EtcdRegistry) serviceNodePath(service *Register.Service) string {
	nodeIp := fmt.Sprintf("%s:%d", service.Nodes[0].Ip, service.Nodes[0].Port)
	return path.Join(e.options.RegistryPath[0], service.Name, nodeIp)
}
