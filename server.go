package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jayanthkrishna/distributed-cache-golang/cache"
)

type ServerOpts struct {
	ListenAddr string
	IsLeader   bool
}

type Server struct {
	ServerOpts

	cache cache.Cacher
}

func NewServer(opts ServerOpts, c cache.Cacher) *Server {
	return &Server{
		ServerOpts: opts,
		cache:      c,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)

	if err != nil {
		return fmt.Errorf("Listen error : %s ", err)
	}

	log.Printf("Server starting at port [%s]\n ", s.ListenAddr)

	for {
		conn, err := ln.Accept()

		if err != nil {
			log.Printf("Acce[t error : %s\n", err)
			continue
		}
		go s.handleConn(conn)

	}
}

func (s *Server) handleConn(conn net.Conn) {

}
