package loadbalance

import (
	"Microservices_demo/MicroService/Register"
	"context"
	"fmt"
	"testing"
)

func TestRandomBalance_Select(t *testing.T) {
	balance := &RandomBalance{}
	var nodes []*Register.Node

	for i := 0; i < 10; i++ {
		node := &Register.Node{
			Ip:   fmt.Sprintf("127.0.0.%d", i),
			Port: 9000 + i,
		}
		nodes = append(nodes, node)
	}

	countStat := make(map[string]int)
	for i := 0; i < 1000; i++ {
		node, err := balance.Select(context.TODO(), nodes)
		if err != nil {
			t.Fatalf("select failed err: %v", err)
			continue
		}

		countStat[node.Ip]++
	}

	for key, val := range countStat {
		fmt.Printf("ip: %s count: %v\n", key, val)
	}
}

func TestRandomBalance_Name(t *testing.T) {
	balance := &RandomBalance{}
	name := balance.Name()
	if name != "random" {
		t.Fatalf("name err: %v", name)
	}
}
