package loadbalance

import (
	"Microservices_demo/MicroService/Register"
	"context"
	"math/rand"
)

type RandomBalance struct {
}

func (r *RandomBalance) Name() string {
	return "random"
}

func (r *RandomBalance) Select(ctx context.Context, nodes []*Register.Node) (node *Register.Node, err error) {
	if len(nodes) == 0 {
		err = ErrNotHaveBodes
		return
	}

	index := rand.Intn(len(nodes))
	node = nodes[index]
	return
}
