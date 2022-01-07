// Created By: junmin.ma
// Description: <description>
// Date: 2021-11-29 11:40
package cfg

import "time"

type Server struct {
	cfg Config
}

type Config struct {
	host    string
	port    int
	timeout time.Duration
	maxConn int
}

func New(cfg Config) *Server {
	return &Server{cfg: cfg}
}

func (s *Server) Start() error {
	// todo
	return nil
}
