package etcd

import (
	"Microservices_demo/MicroService/Register"
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

//ETCDCTL_API=3 ./etcdctl get "/micr/etcd/" --prefix

func TestEtcdRegistry_Register(t *testing.T) {

	//初始化
	registryInst, err := Register.InitRegistry(context.TODO(), "etcd",
		Register.WithAddres([]string{"192.168.31.205:2379"}),
		Register.WithTimeOut(time.Second),
		Register.WithRegistryPath([]string{"/micr/etcd/"}),
		Register.WithRegistryHeartBeat(5),
	)

	if err != nil {
		t.Errorf("registry init failed err %v", err)
		return
	}

	//服务
	service := &Register.Service{
		Name: "comment_services",
	}

	service.Nodes = append(service.Nodes,
		&Register.Node{Ip: "127.0.0.1", Port: 8801},
		&Register.Node{Ip: "127.0.0.1", Port: 8802},
	)

	registryInst.Register(context.TODO(), service)

	go func() {
		time.Sleep(10 * time.Second)

		service.Nodes = append(service.Nodes, &Register.Node{
			Ip: "127.0.0.8", Port: 8801,
		}, &Register.Node{
			Ip: "127.0.0.3", Port: 8802,
		})

		registryInst.Register(context.TODO(), service)
	}()

	for {
		service, err = registryInst.GetService(context.TODO(), "comment_services")
		if err != nil {
			t.Errorf("get service failed err: %v", err)
			return
		}
		for _, v := range service.Nodes {
			log.Printf("name %s, port %d, ip %s\n", service.Name, v.Port, v.Ip)
		}
		fmt.Printf("service %#v\n", service)
		time.Sleep(time.Second)

	}

	return
}
