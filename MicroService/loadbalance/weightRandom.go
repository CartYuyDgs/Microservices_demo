package loadbalance

import (
	"Microservices_demo/MicroService/Register"
	"context"
	"math/rand"
)

type WeightRandomBalance struct {
}

func (w *WeightRandomBalance) Name() string {
	return "WeightRandomBalance"
}

func (w *WeightRandomBalance) Select(ctx context.Context, nodes []*Register.Node) (node *Register.Node, err error) {
	if len(nodes) == 0 {
		err = ErrNotHaveBodes
		return
	}

	var totalWeight int
	for _, val := range nodes {
		if val.Weight == 0 {
			val.Weight = DefaultWeight
		}
		totalWeight = totalWeight + val.Weight
	}

	curWeight := rand.Intn(totalWeight)
	curIndex := 0
	for index, val := range nodes {
		curWeight -= val.Weight
		if curWeight < 0 {
			curIndex = index
			break
		}
	}

	if curIndex == 0 {
		err = ErrNotHaveBodes
		return
	}

	node = nodes[curIndex]
	return
}
