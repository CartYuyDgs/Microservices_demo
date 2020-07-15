package middleware

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
)

func TestMiddleware(t *testing.T) {

	m1 := func(next MiddlewareFunc) MiddlewareFunc {
		return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
			fmt.Println("middleware1 start...")
			num := rand.Intn(2)
			if num <= 2 {
				err = fmt.Errorf("this is request is not allow!")
				return
			}
			resp, err = next(ctx, req)
			if err != nil {
				return
			}
			fmt.Println("middleware1 end...")
			return
		}
	}

	m2 := func(next MiddlewareFunc) MiddlewareFunc {
		return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
			fmt.Println("middleware2 start...")
			resp, err = next(ctx, req)
			if err != nil {
				return
			}
			fmt.Println("middleware2 end...")
			return
		}
	}

	outer := func(next MiddlewareFunc) MiddlewareFunc {
		return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
			fmt.Println("outer start...")
			resp, err = next(ctx, req)
			if err != nil {
				return
			}
			fmt.Println("outer end...")
			return
		}
	}

	proc := func(ctx context.Context, req interface{}) (resp interface{}, err error) {
		fmt.Println("proc start")
		fmt.Println("proc end")
		return
	}

	chain := Chain(outer, m1, m2)
	chainFunc := chain(proc)

	resp, err := chainFunc(context.Background(), "text")
	fmt.Printf("resp:%#v, err:%v\n", resp, err)
	fmt.Println("------------end-----------------")

}
