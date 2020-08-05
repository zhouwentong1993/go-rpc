package go_rpc

type Server struct {
	option  *ServerOptions
	service Service
}

func newServer(opts ...ServerOption) *Server {
	s := &Server{
		option: &ServerOptions{},
	}
	for _, o := range opts {
		o(s.option)
	}
	s.service = newService(s.option)
	return s
}

func newService(opts *ServerOptions) Service {
	return &service{
		opts: opts,
	}

}
