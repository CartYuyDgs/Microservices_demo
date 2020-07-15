package main

import "context"

type MiddlewareFunc func(ctx context.Context, req interface{}) (resp interface{}, err error)

type Middleware func(MiddlewareFunc) MiddlewareFunc

var userMiddleware []Middleware

func Chain(outer Middleware, other ...Middleware) Middleware {
	return func(next MiddlewareFunc) MiddlewareFunc {
		for i := len(other) - 1; i >= 0; i-- {
			next = other[i](next)
		}
		return outer(next)
	}
}

func Use(m ...Middleware) {
	userMiddleware = append(userMiddleware, m...)
}
