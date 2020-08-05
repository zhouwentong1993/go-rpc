package go_rpc

import (
	"context"
	"fmt"
)

// 一个服务能提供的能力
type Service interface {
	Register(string, Handler)
	Serve(*ServerOptions)
	Close()
}

type service struct {
	svr         interface{}
	ctx         context.Context
	cancel      context.CancelFunc
	serviceName string
	handlers    map[string]Handler
	opts        *ServerOptions

	closing bool
}

func (s service) Register(handlerName string, handler Handler) {
	if s.handlers == nil {
		s.handlers = make(map[string]Handler)
	}
	s.handlers[handlerName] = handler
}

func (s service) Serve(opts *ServerOptions) {

}

func (s service) Close() {
	s.closing = true

	if s.cancel != nil {
		s.cancel()
	}
	fmt.Println("closing")
}

type Handler func(context.Context, interface{}, func(interface{}) error) (interface{}, error)
