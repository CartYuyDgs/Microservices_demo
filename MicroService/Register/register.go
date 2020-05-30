package Register

import "context"

type Registry interface {
	Name() string
	Init(ctx context.Context, opts ...Option) (err error)
	Register(ctx context.Context, service *Service) (err error)
	Unregister(ctx context.Context, service *Service) (err error)
	GetService(ctx context.Context, name string) (service *Service, err error)
}

//func Init(opts ...Option){
//
//}
//
//func (*Registry) RegisterService(ctx context.Context, service *Service) {
//
//}
//
//func (*Registry) Unregister(ctx context.Context, service *Service){
//
//}
