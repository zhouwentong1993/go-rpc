package go_rpc

import (
	"time"
)

type Server struct {
	option   *ServerOption
	services map[string]Service
}

type ServerOption struct {
	//地址
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

type service struct {
}

type Service interface {
	Register()
}
