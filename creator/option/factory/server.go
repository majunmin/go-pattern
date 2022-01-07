// Created By: junmin.ma
// Description: <description>
// Date: 2021-11-29 11:36
package factory

import "time"

type Server struct {
	host    string
	port    int
	timeout time.Duration
	maxConn int
}

func New(host string, port int) *Server {
	return &Server{host, port, time.Minute, 100}
}

func NewWithTimeout(host string, port int, timeout time.Duration) *Server {
	return &Server{host, port, timeout, 100}
}

func NewWithTimeoutAndMaxConn(host string, port int, timeout time.Duration, maxConn int) *Server {
	return &Server{host, port, timeout, maxConn}
}

func (s *Server) Start() error {
	// todo
	return nil
}
