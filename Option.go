package go_rpc

import "time"

type ServerOptions struct {
	address string
	//tcp or udp
	network string
	//协议格式
	protocol string
	//超时时间
	timout time.Duration
	// 序列化类型
	serializationType string
	// 服务发现地址
	selectorSvrAddr string // service discovery server address, required when using the third-party service discovery plugin
	// trace 地址
	tracingSvrAddr string // tracing plugin server address, required when using the third-party tracing plugin
	// trace span 名
	tracingSpanName string // tracing span name, required when using the third-party tracing plugin
	// 用的插件
	pluginNames []string // plugin name
}

// option 模式
type ServerOption func(*ServerOptions)

func WithAddress(address string) ServerOption {
	return func(o *ServerOptions) {
		o.address = address
	}
}
func WithNetwork(network string) ServerOption {
	return func(options *ServerOptions) {
		options.network = network
	}
}
func WithProtocol(protocol string) ServerOption {
	return func(o *ServerOptions) {
		o.protocol = protocol
	}
}

func WithSerializationType(serializationType string) ServerOption {
	return func(o *ServerOptions) {
		o.serializationType = serializationType
	}
}

func WithSelectorSvrAddr(addr string) ServerOption {
	return func(o *ServerOptions) {
		o.selectorSvrAddr = addr
	}
}

func WithPlugin(pluginName ...string) ServerOption {
	return func(o *ServerOptions) {
		o.pluginNames = append(o.pluginNames, pluginName...)
	}
}

func WithTracingSvrAddr(addr string) ServerOption {
	return func(o *ServerOptions) {
		o.tracingSvrAddr = addr
	}
}

func WithTracingSpanName(name string) ServerOption {
	return func(o *ServerOptions) {
		o.tracingSpanName = name
	}
}
