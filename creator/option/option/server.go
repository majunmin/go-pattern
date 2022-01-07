// Created By: junmin.ma
// Description: <description>
// Date: 2021-11-29 11:47
package option

import "time"

type Server struct {
	host    string
	port    int
	timeout time.Duration
	maxConn int
}
type Option func(server *Server)

func New(options ...Option) *Server {
	srv := &Server{}
	for _, opt := range options {
		opt(srv)
	}
	return srv
}

func (s *Server) Start() error {
	//todo
	return nil
}

func WithHost(host string) Option {
	return func(s *Server) {
		s.host = host
	}
}

func WithPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithMaxConn(maxConn int) Option {
	return func(s *Server) {
		s.maxConn = maxConn
	}
}
