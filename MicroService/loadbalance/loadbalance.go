package loadbalance

import (
	"Microservices_demo/MicroService/Register"
	"context"
	"errors"
)

var ErrNotHaveBodes = errors.New("not have node!")

const DefaultWeight = 100

type LoadBalance interface {
	Name()
	Select(ctx context.Context, nodes []*Register.Node) (node *Register.Node, err error)
}
